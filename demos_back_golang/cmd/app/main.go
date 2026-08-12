package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"demos_back_golang/internal/config"
	"demos_back_golang/internal/handlers"
	"demos_back_golang/internal/lib/jwt"
	"demos_back_golang/internal/lib/slogpretty/handlers/slogpretty"
	"demos_back_golang/internal/lib/slogpretty/sl"
	"demos_back_golang/internal/middleware"
	"demos_back_golang/internal/storage"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Config
	configPath := flag.String("config", "./configs/local.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatal("Error while loading config: ", err) // Посмотреть реализацию функции MustLoad config в видео, как там логируется инфа
	}

	// Logger
	logger := setupLogger(cfg.App.Env)
	logger.Debug("Logger is working correctly")

	logger.Debug("Config loaded:", "Config", cfg)

	// Database & Storage
	db, err := storage.NewDatabase(cfg.Database)
	if err != nil {
		logger.Error("Failed to connect to database", sl.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	store := storage.NewStorage(db)

	// JWT Service
	jwtService := jwt.NewJwtService(
		cfg.App.JWTService.JwtSecret,
		cfg.App.JWTService.AccessTTL,
		cfg.App.JWTService.RefreshTTL,
	)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	// Dependencies

	userHandler := handlers.NewUserHandler(
		store.Users,
		store.RefreshTokens,
		logger,
		jwtService,
	)
	classHandler := handlers.NewClassHandler(store.Classes, logger)

	// App
	app := &application{
		logger:         logger,
		config:         cfg,
		storage:        store,
		userHandler:    userHandler,
		classHandler:   classHandler,
		authMiddleware: authMiddleware,
	}

	mux := app.mount()

	err = app.run(mux)
	if err != nil {
		logger.Error("Failed to run application", sl.Err(err))
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

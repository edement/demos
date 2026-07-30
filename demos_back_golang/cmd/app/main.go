package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"demos_back_golang/internal/config"
	"demos_back_golang/internal/database"
	"demos_back_golang/internal/lib/slogpretty/handlers/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	configPath := flag.String("config", "./configs/local.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatal("Error while loading config: ", err) // Посмотреть реализацию функции MustLoad config в видео, как там логируется инфа
	}
	log.Println(cfg)

	logger := setupLogger(cfg.App.Environment)
	logger.Info("Logger is working correctly")

	db, err := database.NewDatabase(cfg.Database.Addres)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
	}
	defer db.Close()

	storage := database.NewStorage(db)

	app := &application{
		logger:  logger,
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()

	err = app.run(mux)
	if err != nil {
		logger.Error("Failed to run application %w", "error", err)
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

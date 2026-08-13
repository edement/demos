package main

import (
	"log/slog"
	"net/http"
	"time"

	"demos_back_golang/internal/config"
	"demos_back_golang/internal/handlers"
	"demos_back_golang/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type application struct {
	logger         *slog.Logger
	config         *config.Config
	storage        *storage.Storage
	userHandler    *handlers.UserHandler
	classHandler   *handlers.ClassHandler
	authMiddleware func(http.Handler) http.Handler
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           86400,
	}).Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() yhat the request has timed out and further
	// processing should be stoped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.HealthCheck)

		r.Route("/users", func(r chi.Router) {
			r.Post("/register", app.userHandler.Register)
			r.Post("/login", app.userHandler.Login)
			r.Post("/refresh", app.userHandler.Refresh)

			r.Group(func(r chi.Router) {
				r.Use(app.authMiddleware)
				r.Get("/me", app.userHandler.Me)
			})
		})

		r.Route("/classes", func(r chi.Router) {
			r.Get("/{classId}", app.classHandler.GetClassById)
			r.Get("/", app.classHandler.GetClasses)

			r.Group(func(r chi.Router) {
				r.Use(app.authMiddleware)
				r.Post("/create", app.classHandler.CreateClass)
				r.Delete("/{classId}", app.classHandler.DeleteClass)
				r.Patch("/{classId}", app.classHandler.UpdateClass)
			})
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.Server.Host + ":" + app.config.Server.Port,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	startInfo := "Server starts on " + srv.Addr
	app.logger.Info(startInfo)

	return srv.ListenAndServe()
}

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
)

type application struct {
	logger       *slog.Logger
	config       *config.Config
	storage      *storage.Storage
	userHandler  *handlers.UserHandler
	classHandler *handlers.ClassHandler
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

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
		})

		r.Route("/classes", func(r chi.Router) {
			r.Post("/create", app.classHandler.CreateClass)
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

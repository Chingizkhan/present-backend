package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"present/present/config"
)

func NewRouter(cfg *config.Config, l *slog.Logger) http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("present-backend", map[string]string{
			cfg.HTTP.User: cfg.HTTP.Password,
		}))
		//r.Post("/", save.New(log, storage))
	})

	//router.Get("/{alias}", redirect.New(log, storage))

	return router
}

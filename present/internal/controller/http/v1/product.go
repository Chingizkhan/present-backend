package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"present/present/config"
	"present/present/internal/usecase"
)

type productRoutes struct {
	t usecase.Product
	l *slog.Logger
}

func newProductRoutes(router *chi.Mux, cfg *config.Config, p *usecase.ProductUseCase, l *slog.Logger) {
	r := productRoutes{p, l}

	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("present-backend", map[string]string{
			cfg.HTTP.User: cfg.HTTP.Password,
		}))
		//r.Post("/", save.New(log, storage))
	})

	//router.Get("/{alias}", redirect.New(log, storage))

}

func (r *productRoutes) Save() {
	r.t.Save()
}

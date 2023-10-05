package v1

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"present/present/config"
	"present/present/internal/controller/http/v1/handler/product/find"
	"present/present/internal/controller/http/v1/handler/product/save"
	"present/present/internal/usecase"
)

type productRoutes struct {
	u usecase.Product
	l *slog.Logger
}

func newProductRoutes(router *chi.Mux, cfg *config.Config, u usecase.Product, l *slog.Logger) {
	routes := productRoutes{u, l}

	router.Route("/product", func(r chi.Router) {
		//r.Use(middleware.BasicAuth("present-backend", map[string]string{
		//	cfg.HTTP.User: cfg.HTTP.Password,
		//}))
		r.Post("/add", routes.Save)
		r.Get("/{id}", routes.Find)
		//r.Delete("/delete", save.New(log, storage))
	})
}

func (routes *productRoutes) Save(w http.ResponseWriter, r *http.Request) {
	save.Handle(w, r, routes.l, routes.u)
}

func (routes *productRoutes) Find(w http.ResponseWriter, r *http.Request) {
	find.Handle(w, r, routes.l, routes.u)
}

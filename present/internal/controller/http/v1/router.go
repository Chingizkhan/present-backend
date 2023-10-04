package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"present/present/config"
	mw "present/present/internal/controller/http/v1/middleware"
	"present/present/internal/usecase"
)

func NewRouter(cfg *config.Config, l *slog.Logger, p *usecase.ProductUseCase) http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mw.Logger(l))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	{
		newProductRoutes(router, cfg, p, l)
	}

	return router
}

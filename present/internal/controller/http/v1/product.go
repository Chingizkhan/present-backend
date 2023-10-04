package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"present/present/config"
	"present/present/internal/entity"
	"present/present/internal/usecase"
)

type productRoutes struct {
	t usecase.Product
	l *slog.Logger
}

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Resp
	Alias string `json:"alias,omitempty"`
}

func newProductRoutes(router *chi.Mux, cfg *config.Config, p *usecase.ProductUseCase, l *slog.Logger) {
	routes := productRoutes{p, l}

	router.Route("/product", func(r chi.Router) {
		r.Use(middleware.BasicAuth("present-backend", map[string]string{
			cfg.HTTP.User: cfg.HTTP.Password,
		}))
		r.Post("/add", routes.Save)
		//r.Delete("/delete", save.New(log, storage))
	})
	//router.Get("/{alias}", redirect.New(log, storage))
}

func (routes *productRoutes) Save(w http.ResponseWriter, r *http.Request) {
	const op = "v1 - Save"

	log := routes.l.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req Request

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("failed to decode request body", "error", err)

		render.JSON(w, r, Error("failed to decode request"))

		return
	}

	log.Info("request body decoded", slog.Any("request", req))

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)

		log.Error("invalid request", "error", err)

		render.JSON(w, r, ValidationError(validateErr))

		return
	}

	alias := req.Alias
	if alias == "" {
		alias = ""
	}

	err = routes.t.Save(r.Context(), entity.Produ)
	if err != nil {
		return
	}
}

func responseOk(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Resp:  OK(),
		Alias: alias,
	})
}

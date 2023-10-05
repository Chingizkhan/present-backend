package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	Name  string `json:"name" validate:"required"`
	Brand string `json:"brand" validate:"required"`
	//Category string `json:"category" validate:"required"`
	//Currency string `json:"currency" validate:"required"`
	//Quantity int    `json:"quantity" validate:"required"`
	//Price    int    `json:"price" validate:"required"`
	//OldPrice int    `json:"old_price,omitempty"`
}

type Response struct {
	Data   FindResponseProduct `json:"data"`
	Status Resp                `json:"status"`
}

type FindResponseProduct struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

func newProductRoutes(router *chi.Mux, cfg *config.Config, p *usecase.ProductUseCase, l *slog.Logger) {
	routes := productRoutes{p, l}

	router.Route("/product", func(r chi.Router) {
		//r.Use(middleware.BasicAuth("present-backend", map[string]string{
		//	cfg.HTTP.User: cfg.HTTP.Password,
		//}))
		r.Get("/{id}", routes.Find)
		r.Post("/add", routes.Save)
		//r.Delete("/delete", save.New(log, storage))
	})
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

	err = routes.t.Save(r.Context(), mapProductRequestToEntity(req))
	if err != nil {
		log.Error("can not save product", "error", err)
		render.JSON(w, r, Error("can not save product"))
		return
	}

	render.JSON(w, r, OK())
}

func (routes *productRoutes) Find(w http.ResponseWriter, r *http.Request) {
	const op = "v1 - Find"

	log := routes.l.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	if id == "" {
		log.Info("id is empty")
		render.JSON(w, r, "invalid request")
		return
	}
	uuidId, err := uuid.Parse(id)
	if err != nil {
		log.Info("can not parse id to uuid")
		render.JSON(w, r, "invalid id")
		return
	}

	data, err := routes.t.Find(r.Context(), uuidId)
	if err != nil {
		log.Error("can not find by this id - "+id, "error", err.Error())
		render.JSON(w, r, "can not find by this id - "+id)
		return
	}

	if len(data) == 0 {
		log.Info("products not found")
		render.JSON(w, r, "products not found")
		return
	}

	render.JSON(w, r, Response{
		Data: FindResponseProduct{
			Id:    data[0].Id.String(),
			Name:  data[0].Name,
			Brand: data[0].Brand,
		},
		Status: OK(),
	})
}

func mapProductRequestToEntity(req Request) entity.Product {
	return entity.Product{
		Id:    uuid.New(),
		Name:  req.Name,
		Brand: req.Brand,
		//Category: req.Category,
		//Currency: req.Currency,
		//Quantity: req.Quantity,
		//Price:    req.Price,
		//OldPrice: req.OldPrice,
	}
}

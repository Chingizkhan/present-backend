package save

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"present/present/internal/controller/http/v1/api"
	"present/present/internal/entity"
	"present/present/internal/usecase"
)

type Request struct {
	Name  string `json:"name" validate:"required"`
	Brand string `json:"brand" validate:"required"`
	//Category string `json:"category" validate:"required"`
	//Currency string `json:"currency" validate:"required"`
	//Quantity int    `json:"quantity" validate:"required"`
	//Price    int    `json:"price" validate:"required"`
	//OldPrice int    `json:"old_price,omitempty"`
}

func Handle(w http.ResponseWriter, r *http.Request, l *slog.Logger, u usecase.Product) {
	const op = "v1 - Save"

	l = l.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req Request

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		l.Error("failed to decode request body", "error", err)
		render.JSON(w, r, api.Error("failed to decode request"))
		return
	}
	l.Info("request body decoded", slog.Any("request", req))

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)
		l.Error("invalid request", "error", err)
		render.JSON(w, r, api.ValidationError(validateErr))
		return
	}

	err = u.Save(r.Context(), mapProductRequestToEntity(req))
	if err != nil {
		l.Error("can not save product", "error", err)
		render.JSON(w, r, api.Error("can not save product"))
		return
	}

	render.JSON(w, r, api.OK())
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

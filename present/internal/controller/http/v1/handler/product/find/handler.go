package find

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"present/present/internal/controller/http/v1/api"
	"present/present/internal/usecase"
)

func Handle(w http.ResponseWriter, r *http.Request, l *slog.Logger, u usecase.Product) {
	const op = "v1 - Find"

	log := l.With(
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

	data, err := u.Find(r.Context(), uuidId)
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
		Data: Product{
			Id:    data[0].Id.String(),
			Name:  data[0].Name,
			Brand: data[0].Brand,
		},
		Status: api.OK(),
	})
}

type Response struct {
	Status api.Resp `json:"status"`
	Data   Product  `json:"data"`
}

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

package usecase

import (
	"context"
	"present/present/internal/entity"
)

type (
	Product interface {
		Find(ctx context.Context) ([]entity.Product, error)
		Save(ctx context.Context, p entity.Product) error
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, p entity.Product) error
	}

	ProductRepo interface {
		Find(ctx context.Context) ([]entity.Product, error)
		Save(ctx context.Context, p entity.Product) error
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, p entity.Product) error
	}
)

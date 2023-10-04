package usecase

import (
	"context"
	"fmt"
	"present/present/internal/entity"
)

type ProductUseCase struct {
	repo ProductRepo
}

func New(r ProductRepo) *ProductUseCase {
	return &ProductUseCase{
		repo: r,
	}
}

func (uc *ProductUseCase) Find(ctx context.Context) ([]entity.Product, error) {
	const op = "ProductUseCase - Find"

	products, err := uc.repo.Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s - uc.repo.Find: %w", op, err)
	}

	return products, nil
}

func (uc *ProductUseCase) Save(ctx context.Context, p entity.Product) error {
	return nil
}

func (uc *ProductUseCase) Delete(ctx context.Context, id string) error {
	return nil
}

func (uc *ProductUseCase) Update(ctx context.Context, p entity.Product) error {
	return nil
}

package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func (uc *ProductUseCase) Find(ctx context.Context, id uuid.UUID) ([]entity.Product, error) {
	const op = "ProductUseCase - Find"

	products, err := uc.repo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s - uc.repo.Find: %w", op, err)
	}

	return products, nil
}

func (uc *ProductUseCase) Save(ctx context.Context, p entity.Product) error {
	const op = "ProductUseCase - Save"

	err := uc.repo.Save(ctx, p)
	if err != nil {
		return fmt.Errorf("%s - uc.repo.Save: %w", op, err)
	}

	return nil
}

func (uc *ProductUseCase) Delete(ctx context.Context, id string) error {
	return nil
}

func (uc *ProductUseCase) Update(ctx context.Context, p entity.Product) error {
	return nil
}

package repo

import (
	"context"
	"fmt"
	"present/present/internal/entity"
	"present/present/pkg/postgres"
)

const _defaultEntityCap = 100

type ProductRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{pg}
}

func (r *ProductRepo) Find(ctx context.Context) ([]entity.Product, error) {
	const op = "ProductRepo - Find"

	sql, _, err := r.Builder.
		Select("*").
		From("product").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s - r.Builder: %w", op, err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("%s - r.Pool.Query: %w", op, err)
	}
	defer rows.Close()

	entities := make([]entity.Product, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Product{}

		if err := rows.Scan(&e.Id); err != nil {
			return nil, fmt.Errorf("%s - rows.Scan: %w", op, err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (r *ProductRepo) Save(ctx context.Context, p entity.Product) error {
	return nil
}

func (r *ProductRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *ProductRepo) Update(ctx context.Context, p entity.Product) error {
	return nil
}

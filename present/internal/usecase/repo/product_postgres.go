package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"present/present/internal/entity"
	"present/present/pkg/postgres"
)

const (
	_defaultEntityCap = 100
	tableName         = "product"
)

type ProductRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{pg}
}

func (r *ProductRepo) Find(ctx context.Context, id uuid.UUID) ([]entity.Product, error) {
	const op = "ProductRepo - Find"

	sql, args, err := r.Builder.
		Select("*").
		From(tableName).
		Where("id = ?", id.String()).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s - r.Builder: %w", op, err)
	}

	log.Println("sql:", sql)
	log.Printf("args: %#v", args)

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s - r.Pool.Query: %w", op, err)
	}
	defer rows.Close()

	entities := make([]entity.Product, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Product{}

		if err := rows.Scan(&e.Id, &e.Name, &e.Brand); err != nil {
			return nil, fmt.Errorf("%s - rows.Scan: %w", op, err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (r *ProductRepo) Save(ctx context.Context, p entity.Product) error {
	const op = "ProductRepo - Save"

	sql, args, err := r.Builder.
		Insert(tableName).
		Columns("id, name, brand").
		Values(p.Id, p.Name, p.Brand).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s - r.Builder: %w", op, err)
	}

	log.Println("sql:", sql)
	log.Printf("args: %#v", args)

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("%s - r.Pool.Exec: %w", op, err)
	}

	return nil
}

func (r *ProductRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *ProductRepo) Update(ctx context.Context, p entity.Product) error {
	return nil
}

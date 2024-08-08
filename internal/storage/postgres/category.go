package storage

import (
	"context"
	"github.com/Sanchir01/microservice_sandjma_category/internal/modain/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"time"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (str *Storage) AllCategory(ctx context.Context) ([]models.Category, error) {
	const op = "storage.postgres.AllCategory"

	conn, err := str.db.Connx(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	var category []dbCategory
	if err := conn.SelectContext(ctx, &category, "EXPLAIN ANALYZE SELECT * FROM categories"); err != nil {
		return nil, err
	}

	return lo.Map(category, func(category dbCategory, _ int) models.Category { return models.Category(category) }), nil
}

type dbCategory struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Slug        string    `db:"slug"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Description string    `db:"description"`
	Version     uint      `db:"version"`
}

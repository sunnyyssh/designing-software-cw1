package pgrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/storage"
	"github.com/sunnyyssh/designing-software-cw1/internal/domain"
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) Get(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	query := `
		SELECT id, type, name
		FROM categories
		WHERE id = $1
	`

	var category domain.Category
	err := r.db.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Type,
		&category.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	return &category, nil
}

func (r *CategoryRepo) GetByType(ctx context.Context, categoryType domain.CategoryType) (*domain.Category, error) {
	query := `
		SELECT id, type, name
		FROM categories
		WHERE type = $1
	`

	var category domain.Category
	err := r.db.QueryRow(ctx, query, categoryType).Scan(
		&category.ID,
		&category.Type,
		&category.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get category by type: %w", err)
	}

	return &category, nil
}

func (r *CategoryRepo) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	query := `
		INSERT INTO categories (id, type, name)
		VALUES ($1, $2, $3)
		RETURNING id, type, name
	`

	err := r.db.QueryRow(ctx, query,
		category.ID,
		category.Type,
		category.Name,
	).Scan(
		&category.ID,
		&category.Type,
		&category.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return category, nil
}

func (r *CategoryRepo) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	query := `
		UPDATE categories
		SET type = $2, name = $3
		WHERE id = $1
		RETURNING id, type, name
	`

	err := r.db.QueryRow(ctx, query,
		category.ID,
		category.Type,
		category.Name,
	).Scan(
		&category.ID,
		&category.Type,
		&category.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return category, nil
}

func (r *CategoryRepo) Delete(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	query := `
		DELETE FROM categories
		WHERE id = $1
		RETURNING id, type, name
	`

	var category domain.Category
	err := r.db.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Type,
		&category.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to delete category: %w", err)
	}

	return &category, nil
}

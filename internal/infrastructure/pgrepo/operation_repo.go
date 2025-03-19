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

type OperationRepo struct {
	db *pgxpool.Pool
}

func NewOperationRepo(db *pgxpool.Pool) *OperationRepo {
	return &OperationRepo{db: db}
}

func (r *OperationRepo) Get(ctx context.Context, id uuid.UUID) (*domain.Operation, error) {
	query := `
		SELECT id, account_id, type, amount, time, description, category_id
		FROM operations
		WHERE id = $1
	`

	var operation domain.Operation
	err := r.db.QueryRow(ctx, query, id).Scan(
		&operation.ID,
		&operation.AccountID,
		&operation.Type,
		&operation.Amount,
		&operation.Time,
		&operation.Description,
		&operation.CategoryID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get operation: %w", err)
	}

	return &operation, nil
}

func (r *OperationRepo) Create(ctx context.Context, operation *domain.Operation) (*domain.Operation, error) {
	query := `
		INSERT INTO operations (id, account_id, type, amount, time, description, category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, account_id, type, amount, time, description, category_id
	`

	err := r.db.QueryRow(ctx, query,
		operation.ID,
		operation.AccountID,
		operation.Type,
		operation.Amount,
		operation.Time,
		operation.Description,
		operation.CategoryID,
	).Scan(
		&operation.ID,
		&operation.AccountID,
		&operation.Type,
		&operation.Amount,
		&operation.Time,
		&operation.Description,
		&operation.CategoryID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create operation: %w", err)
	}

	return operation, nil
}

func (r *OperationRepo) Update(ctx context.Context, operation *domain.Operation) (*domain.Operation, error) {
	query := `
		UPDATE operations
		SET account_id = $2, type = $3, amount = $4, time = $5, description = $6, category_id = $7
		WHERE id = $1
		RETURNING id, account_id, type, amount, time, description, category_id
	`

	err := r.db.QueryRow(ctx, query,
		operation.ID,
		operation.AccountID,
		operation.Type,
		operation.Amount,
		operation.Time,
		operation.Description,
		operation.CategoryID,
	).Scan(
		&operation.ID,
		&operation.AccountID,
		&operation.Type,
		&operation.Amount,
		&operation.Time,
		&operation.Description,
		&operation.CategoryID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to update operation: %w", err)
	}

	return operation, nil
}

func (r *OperationRepo) Delete(ctx context.Context, id uuid.UUID) (*domain.Operation, error) {
	query := `
		DELETE FROM operations
		WHERE id = $1
		RETURNING id, account_id, type, amount, time, description, category_id
	`

	var operation domain.Operation
	err := r.db.QueryRow(ctx, query, id).Scan(
		&operation.ID,
		&operation.AccountID,
		&operation.Type,
		&operation.Amount,
		&operation.Time,
		&operation.Description,
		&operation.CategoryID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to delete operation: %w", err)
	}

	return &operation, nil
}

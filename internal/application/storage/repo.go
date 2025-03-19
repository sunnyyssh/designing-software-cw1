package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sunnyyssh/designing-software-cw1/internal/domain"
)

var (
	ErrNotFound = errors.New("not found")
)

type BankAccountRepo interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.BankAccount, error)
	Update(context.Context, *domain.BankAccount) (*domain.BankAccount, error)
	Create(context.Context, *domain.BankAccount) (*domain.BankAccount, error)
	Delete(ctx context.Context, id uuid.UUID) (*domain.BankAccount, error)
}

type CategoryRepo interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Category, error)
	GetByType(context.Context, domain.CategoryType) (*domain.Category, error)
	Update(context.Context, *domain.Category) (*domain.Category, error)
	Create(context.Context, *domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, id uuid.UUID) (*domain.Category, error)
}

type OperationRepo interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Operation, error)
	Update(context.Context, *domain.Operation) (*domain.Operation, error)
	Create(context.Context, *domain.Operation) (*domain.Operation, error)
	Delete(ctx context.Context, id uuid.UUID) (*domain.Operation, error)
}

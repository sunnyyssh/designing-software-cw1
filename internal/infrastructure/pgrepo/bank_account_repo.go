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

type BankAccountRepo struct {
	db *pgxpool.Pool
}

func NewBankAccountRepo(db *pgxpool.Pool) *BankAccountRepo {
	return &BankAccountRepo{
		db: db,
	}
}
func (r *BankAccountRepo) Get(ctx context.Context, id uuid.UUID) (*domain.BankAccount, error) {
	query := `
		SELECT id, name, balance, blocked
		FROM bank_accounts
		WHERE id = $1
	`

	var account domain.BankAccount
	err := r.db.QueryRow(ctx, query, id).Scan(
		&account.ID,
		&account.Name,
		&account.Balance,
		&account.Blocked,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get bank account: %w", err)
	}

	return &account, nil
}

func (r *BankAccountRepo) Create(ctx context.Context, account *domain.BankAccount) (*domain.BankAccount, error) {
	query := `
		INSERT INTO bank_accounts (id, name, balance, blocked)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, balance, blocked
	`

	err := r.db.QueryRow(ctx, query,
		account.ID,
		account.Name,
		account.Balance,
		account.Blocked,
	).Scan(
		&account.ID,
		&account.Name,
		&account.Balance,
		&account.Blocked,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create bank account: %w", err)
	}

	return account, nil
}

func (r *BankAccountRepo) Update(ctx context.Context, account *domain.BankAccount) (*domain.BankAccount, error) {
	query := `
		UPDATE bank_accounts
		SET name = $2, balance = $3, blocked = $4
		WHERE id = $1
		RETURNING id, name, balance, blocked
	`

	err := r.db.QueryRow(ctx, query,
		account.ID,
		account.Name,
		account.Balance,
		account.Blocked,
	).Scan(
		&account.ID,
		&account.Name,
		&account.Balance,
		&account.Blocked,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to update bank account: %w", err)
	}

	return account, nil
}

func (r *BankAccountRepo) Delete(ctx context.Context, id uuid.UUID) (*domain.BankAccount, error) {
	query := `
		DELETE FROM bank_accounts
		WHERE id = $1
		RETURNING id, name, balance, blocked
	`

	var account domain.BankAccount
	err := r.db.QueryRow(ctx, query, id).Scan(
		&account.ID,
		&account.Name,
		&account.Balance,
		&account.Blocked,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to delete bank account: %w", err)
	}

	return &account, nil
}

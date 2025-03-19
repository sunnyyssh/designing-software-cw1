package config

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/storage"
	"github.com/sunnyyssh/designing-software-cw1/internal/infrastructure/pgrepo"
)

type DB struct {
	BankAccountRepo storage.BankAccountRepo
	CategoryRepo    storage.CategoryRepo
	OperationRepo   storage.OperationRepo
}

func NewDB(db *pgxpool.Pool) *DB {
	return &DB{
		BankAccountRepo: pgrepo.NewBankAccountRepo(db),
		CategoryRepo:    pgrepo.NewCategoryRepo(db),
		OperationRepo:   pgrepo.NewOperationRepo(db),
	}
}

package config

import (
	"github.com/jackc/pgx/v5"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/storage"
)

type DB struct {
	BankAccountRepo storage.BankAccountRepo
	CategoryRepo    storage.CategoryRepo
	OperationRepo   storage.OperationRepo
}

func NewDB(db *pgx.Conn) *DB {
	return &DB{}
}

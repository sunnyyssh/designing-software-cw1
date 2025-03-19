package domain

import (
	"github.com/google/uuid"
)

type BankAccount struct {
	ID      uuid.UUID
	Name    string
	Balance int64
	Blocked bool
}

func NewBankAccount(name string) (*BankAccount, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	return &BankAccount{
		ID:      uuid.New(), // It should be set in the database creation
		Name:    name,
		Balance: 0,
		Blocked: false,
	}, nil
}

func (a *BankAccount) Block() error {
	if a.Blocked {
		return ErrAlreadyBlocked
	}
	a.Blocked = true
	return nil
}

func (a *BankAccount) Unblock() error {
	if !a.Blocked {
		return ErrAlreadyUnblocked
	}
	a.Blocked = false
	return nil
}

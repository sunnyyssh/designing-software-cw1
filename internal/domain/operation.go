package domain

import (
	"time"

	"github.com/google/uuid"
)

// For testability
var TimeFunc = func() time.Time { return time.Now() }

type OperationType string

const (
	OperationTypeIncome  OperationType = "income"
	OperationTypeOutcome OperationType = "outcome"
)

type Operation struct {
	ID        uuid.UUID
	AccountID uuid.UUID

	Type        OperationType
	Amount      int64
	Time        time.Time
	Description string
	CategoryID  *uuid.UUID
	applied     bool
}

func newOperation(
	accID uuid.UUID,
	typ OperationType,
	amount int64,
	description string,
) (*Operation, error) {
	return &Operation{
		ID:          uuid.New(), // Should be set in DB
		AccountID:   accID,
		Type:        typ,
		Amount:      amount,
		Time:        TimeFunc(),
		Description: description,
		CategoryID:  nil,
	}, nil
}

func (o *Operation) apply(acc *BankAccount) error {
	if o.applied {
		return ErrAlreadyApplied
	}
	if acc.Blocked {
		return ErrAccountBlocked
	}

	switch o.Type {
	case OperationTypeIncome:
		acc.Balance += o.Amount
	case OperationTypeOutcome:
		if acc.Balance-o.Amount < 0 {
			return ErrNotEnoughMoney
		}
		acc.Balance -= o.Amount
	default:
		panic("idk what this op type means")
	}

	o.applied = true
	return nil
}

func (o *Operation) SetCategory(cat *Category) error {
	o.CategoryID = &cat.ID
	return nil
}

func ApplyOperation(
	acc *BankAccount,
	typ OperationType,
	amount int64,
	description string,
) (*Operation, error) {
	op, err := newOperation(acc.ID, typ, amount, description)
	if err != nil {
		return nil, err
	}

	err = op.apply(acc)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func ResolveCategoryType(op *Operation) (CategoryType, error) {
	switch op.Type {
	case OperationTypeIncome:
		return CategoryTypeIncome, nil
	case OperationTypeOutcome:
		return CategoryTypeOutcome, nil
	default:
		return "", ErrCannotResolveCategory
	}
}

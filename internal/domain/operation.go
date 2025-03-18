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
	CategoryID  uuid.UUID
}

func NewOperation(
	accID uuid.UUID,
	typ OperationType,
	amount int64,
	description string,
	categoryID uuid.UUID,
) (*Operation, error) {
	return &Operation{
		ID:          uuid.Nil, // Should be set in DB
		AccountID:   accID,
		Type:        typ,
		Amount:      amount,
		Time:        TimeFunc(),
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

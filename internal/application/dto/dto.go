package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/sunnyyssh/designing-software-cw1/internal/domain"
)

type BankAccountDTO struct {
	ID      uuid.UUID
	Name    string
	Balance int64
	Blocked bool
}

func NewBankAccountDTO(dom *domain.BankAccount) *BankAccountDTO {
	if dom == nil {
		return nil
	}
	return &BankAccountDTO{
		ID:      dom.ID,
		Name:    dom.Name,
		Balance: dom.Balance,
		Blocked: dom.Blocked,
	}
}

type CategoryDTO struct {
	ID   uuid.UUID `json:"id"`
	Type string    `json:"type"`
	Name string    `json:"name"`
}

func NewCategoryDTO(category *domain.Category) *CategoryDTO {
	if category == nil {
		return nil
	}
	return &CategoryDTO{
		ID:   category.ID,
		Type: string(category.Type),
		Name: category.Name,
	}
}

type OperationDTO struct {
	ID          uuid.UUID  `json:"id"`
	AccountID   uuid.UUID  `json:"account_id"`
	Type        string     `json:"type"`
	Amount      int64      `json:"amount"`
	Time        time.Time  `json:"time"`
	Description string     `json:"description"`
	CategoryID  *uuid.UUID `json:"category_id"`
}

func NewOperationDTO(operation *domain.Operation) *OperationDTO {
	if operation == nil {
		return nil
	}
	return &OperationDTO{
		ID:          operation.ID,
		AccountID:   operation.AccountID,
		Type:        string(operation.Type),
		Amount:      operation.Amount,
		Time:        operation.Time,
		Description: operation.Description,
		CategoryID:  operation.CategoryID,
	}
}

package domain

import "github.com/google/uuid"

type CategoryType string

const (
	CategoryTypeIncome  CategoryType = "income"
	CategoryTypeOutcome CategoryType = "outcome"
)

type Category struct {
	ID   uuid.UUID
	Type CategoryType
	Name string
}

func NewCategory(typ CategoryType, name string) (*Category, error) {
	return &Category{
		ID:   uuid.Nil,
		Type: typ,
		Name: name,
	}, nil
}

package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/dto"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/storage"
	"github.com/sunnyyssh/designing-software-cw1/internal/domain"
)

type CategoryService struct {
	catRepo storage.CategoryRepo
}

func NewCategoryService(catRepo storage.CategoryRepo) *CategoryService {
	return &CategoryService{
		catRepo: catRepo,
	}
}

func (s *CategoryService) Create(ctx context.Context, typ string, name string) (*dto.CategoryDTO, error) {
	category, err := domain.NewCategory(domain.CategoryType(typ), name)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	category, err = s.catRepo.Create(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to save category: %w", err)
	}

	return dto.NewCategoryDTO(category), nil
}

func (s *CategoryService) Get(ctx context.Context, id uuid.UUID) (*dto.CategoryDTO, error) {
	category, err := s.catRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	return dto.NewCategoryDTO(category), nil
}

func (s *CategoryService) List(ctx context.Context) ([]dto.CategoryDTO, error) {
	cats, err := s.catRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.CategoryDTO, 0, len(cats))
	for _, c := range cats {
		resp = append(resp, *dto.NewCategoryDTO(&c))
	}
	return resp, nil
}

func (s *CategoryService) Delete(ctx context.Context, id uuid.UUID) (*dto.CategoryDTO, error) {
	category, err := s.catRepo.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete category: %w", err)
	}

	return dto.NewCategoryDTO(category), nil
}

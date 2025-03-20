package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/dto"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/storage"
	"github.com/sunnyyssh/designing-software-cw1/internal/domain"
)

type BankAccountService struct {
	accRepo storage.BankAccountRepo
}

func NewBankAccountService(repo storage.BankAccountRepo) *BankAccountService {
	return &BankAccountService{
		accRepo: repo,
	}
}

func (s *BankAccountService) Get(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error) {
	acc, err := s.accRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewBankAccountDTO(acc), nil
}

func (s *BankAccountService) List(ctx context.Context) ([]dto.BankAccountDTO, error) {
	cats, err := s.accRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.BankAccountDTO, 0, len(cats))
	for _, c := range cats {
		resp = append(resp, *dto.NewBankAccountDTO(&c))
	}
	return resp, nil
}

func (s *BankAccountService) CreateAccount(ctx context.Context, name string) (*dto.BankAccountDTO, error) {
	acc, err := domain.NewBankAccount(name)
	if err != nil {
		return nil, err
	}

	acc, err = s.accRepo.Create(ctx, acc)

	return dto.NewBankAccountDTO(acc), err
}

func (s *BankAccountService) Block(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error) {
	acc, err := s.accRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := acc.Block(); err != nil {
		return nil, err
	}

	if acc, err = s.accRepo.Update(ctx, acc); err != nil {
		return nil, err
	}
	return dto.NewBankAccountDTO(acc), err
}

func (s *BankAccountService) Unblock(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error) {
	acc, err := s.accRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := acc.Unblock(); err != nil {
		return nil, err
	}

	if acc, err = s.accRepo.Update(ctx, acc); err != nil {
		return nil, err
	}
	return dto.NewBankAccountDTO(acc), err
}

func (s *BankAccountService) Delete(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error) {
	acc, err := s.accRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := acc.Delete(); err != nil {
		return nil, err
	}

	if acc, err = s.accRepo.Delete(ctx, acc.ID); err != nil {
		return nil, err
	}
	return dto.NewBankAccountDTO(acc), err
}

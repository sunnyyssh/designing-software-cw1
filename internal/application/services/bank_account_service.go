package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/dto"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/storage"
	"github.com/sunnyyssh/designing-software-cw1/internal/domain"
)

type BankService struct {
	accRepo storage.BankAccountRepo
	catRepo storage.CategoryRepo
	opRepo  storage.OperationRepo
}

func NewBankService(
	repo storage.BankAccountRepo,
	catRepo storage.CategoryRepo,
	opRepo storage.OperationRepo,
) *BankService {
	return &BankService{
		accRepo: repo,
		catRepo: catRepo,
		opRepo:  opRepo,
	}
}

func (s *BankService) CreateAccount(ctx context.Context, name string) (*dto.BankAccountDTO, error) {
	acc, err := domain.NewBankAccount(name)
	if err != nil {
		return nil, err
	}

	acc, err = s.accRepo.Create(ctx, acc)

	return dto.NewBankAccountDTO(acc), err
}

type ApplyOperationRequest struct {
	AccountID     uuid.UUID
	Amount        int64
	OperationType string
}

func (s *BankService) ApplyOperation(ctx context.Context, req ApplyOperationRequest) (*dto.BankAccountDTO, error) {
	acc, err := s.accRepo.Get(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	op, err := domain.ApplyOperation(acc, domain.OperationType(req.OperationType), req.Amount, "")
	if err != nil {
		return nil, err
	}

	catType, err := domain.ResolveCategoryType(op)
	if err != nil {
		return nil, err
	}

	cat, err := s.catRepo.GetByType(ctx, catType)
	if err != nil {
		return nil, err
	}
	if err := op.SetCategory(cat); err != nil {
		return nil, err
	}

	acc, err = s.accRepo.Update(ctx, acc)
	if err != nil {
		return nil, err
	}
	_, err = s.opRepo.Create(ctx, op)
	if err != nil {
		return nil, err
	}
	return dto.NewBankAccountDTO(acc), nil
}

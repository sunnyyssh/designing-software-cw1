package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/dto"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/storage"
	"github.com/sunnyyssh/designing-software-cw1/internal/domain"
)

type OperationService struct {
	accRepo storage.BankAccountRepo
	opRepo  storage.OperationRepo
}

func NewOperationService(
	repo storage.BankAccountRepo,
	opRepo storage.OperationRepo,
) *OperationService {
	return &OperationService{
		accRepo: repo,
		opRepo:  opRepo,
	}
}

func (s *OperationService) Get(ctx context.Context, id uuid.UUID) (*dto.OperationDTO, error) {
	op, err := s.opRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewOperationDTO(op), nil
}

func (s *OperationService) List(ctx context.Context) ([]dto.OperationDTO, error) {
	cats, err := s.opRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.OperationDTO, 0, len(cats))
	for _, c := range cats {
		resp = append(resp, *dto.NewOperationDTO(&c))
	}
	return resp, nil
}

type ApplyOperationRequest struct {
	AccountID     uuid.UUID
	Amount        int64
	OperationType string
}

func (s *OperationService) ApplyOperation(ctx context.Context, req ApplyOperationRequest) (*dto.BankAccountDTO, error) {
	acc, err := s.accRepo.Get(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	op, err := domain.ApplyOperation(acc, domain.OperationType(req.OperationType), req.Amount, "")
	if err != nil {
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

type TransferRequest struct {
	FromAccountID uuid.UUID
	ToAccountID   uuid.UUID
	Amount        int64
}

type TransferResponse struct {
	FromAccount *dto.BankAccountDTO
	ToAccount   *dto.BankAccountDTO
}

func (s *OperationService) Transfer(ctx context.Context, req TransferRequest) (*TransferResponse, error) {
	from, err := s.accRepo.Get(ctx, req.FromAccountID)
	if err != nil {
		return nil, err
	}
	to, err := s.accRepo.Get(ctx, req.ToAccountID)
	if err != nil {
		return nil, err
	}

	opFrom, err := domain.ApplyOperation(from, domain.OperationTypeOutcome, req.Amount, "")
	if err != nil {
		return nil, err
	}
	opTo, err := domain.ApplyOperation(to, domain.OperationTypeIncome, req.Amount, "")
	if err != nil {
		return nil, err
	}

	if from, err = s.accRepo.Update(ctx, from); err != nil {
		return nil, err
	}
	if from, err = s.accRepo.Update(ctx, from); err != nil {
		return nil, err
	}
	if _, err = s.opRepo.Create(ctx, opFrom); err != nil {
		return nil, err
	}
	if _, err = s.opRepo.Create(ctx, opTo); err != nil {
		return nil, err
	}

	return &TransferResponse{
		FromAccount: dto.NewBankAccountDTO(from),
		ToAccount:   dto.NewBankAccountDTO(to),
	}, nil
}

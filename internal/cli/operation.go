package cli

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/dto"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/services"
)

type OperationService interface {
	Get(ctx context.Context, id uuid.UUID) (*dto.OperationDTO, error)
	List(ctx context.Context) ([]dto.OperationDTO, error)
	ApplyOperation(ctx context.Context, req services.ApplyOperationRequest) (*dto.BankAccountDTO, error)
	Transfer(ctx context.Context, req services.TransferRequest) (*services.TransferResponse, error)
}

func Operation(svc OperationService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "operation",
		Short: "Operation-connected actions",
	}
	cmd.AddCommand(
		getOperation(svc),
		listOperations(svc),
		applyIncome(svc),
		applyOutcome(svc),
		transfer(svc),
	)
	return cmd
}

func getOperation(svc OperationService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get an operation by its ID",
	}

	var operationIDStr string
	cmd.Flags().StringVarP(&operationIDStr, "id", "i", "", "Operation ID")
	cmd.MarkFlagRequired("id")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		operationID, err := uuid.Parse(operationIDStr)
		if err != nil {
			return fmt.Errorf("invalid operation ID: %w", err)
		}

		operation, err := svc.Get(cmd.Context(), operationID)
		if err != nil {
			return fmt.Errorf("failed to get operation: %w", err)
		}

		cmd.Println("Operation details:")
		PrettyJSON(cmd, operation)
		return nil
	}

	return cmd
}

func listOperations(svc OperationService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all operations",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		operations, err := svc.List(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list operations: %w", err)
		}

		cmd.Println("Operations:")
		PrettyJSON(cmd, operations)
		return nil
	}

	return cmd
}

func applyIncome(svc OperationService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "income",
		Short: "Apply income operation on account",
	}

	var (
		accIDstr string
		amount   int64
	)
	cmd.PersistentFlags().StringVarP(&accIDstr, "acc-id", "i", "", "Account ID")
	cmd.PersistentFlags().Int64VarP(&amount, "amount", "m", 0, "Amount of money")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		accID, err := uuid.Parse(accIDstr)
		if err != nil {
			return err
		}

		acc, err := svc.ApplyOperation(cmd.Context(), services.ApplyOperationRequest{
			AccountID:     accID,
			Amount:        amount,
			OperationType: "income",
		})
		if err != nil {
			return err
		}

		cmd.Println(`Income operation applied on account`)
		PrettyJSON(cmd, acc)
		return nil
	}
	return cmd
}

func applyOutcome(svc OperationService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "outcome",
		Short: "Apply outcome operation on account",
	}

	var (
		accIDstr string
		amount   int64
	)
	cmd.PersistentFlags().StringVarP(&accIDstr, "acc-id", "i", "", "Account ID")
	cmd.PersistentFlags().Int64VarP(&amount, "amount", "m", 0, "Amount of money")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		accID, err := uuid.Parse(accIDstr)
		if err != nil {
			return err
		}

		acc, err := svc.ApplyOperation(cmd.Context(), services.ApplyOperationRequest{
			AccountID:     accID,
			Amount:        amount,
			OperationType: "outcome",
		})
		if err != nil {
			return err
		}

		cmd.Println(`Outcome operation applied on account`)
		PrettyJSON(cmd, acc)
		return nil
	}
	return cmd
}

func transfer(svc OperationService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Transfer money from one account to the another one",
	}

	var (
		fromAccIDstr string
		toAccIDstr   string
		amount       int64
	)
	cmd.PersistentFlags().StringVarP(&fromAccIDstr, "from-acc-id", "f", "", "From account ID")
	cmd.PersistentFlags().StringVarP(&toAccIDstr, "to-acc-id", "t", "", "To account ID")
	cmd.PersistentFlags().Int64VarP(&amount, "amount", "m", 0, "Amount of money")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		fromAccID, err := uuid.Parse(fromAccIDstr)
		if err != nil {
			return err
		}
		toAccID, err := uuid.Parse(toAccIDstr)
		if err != nil {
			return err
		}

		resp, err := svc.Transfer(cmd.Context(), services.TransferRequest{
			FromAccountID: fromAccID,
			ToAccountID:   toAccID,
			Amount:        amount,
		})
		if err != nil {
			return err
		}

		cmd.Println(`Outcome operation applied on account`)
		PrettyJSON(cmd, resp.FromAccount)
		cmd.Println(`Income operation applied on account`)
		PrettyJSON(cmd, resp.ToAccount)
		return nil
	}
	return cmd
}

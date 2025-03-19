package cli

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/services"
	"github.com/sunnyyssh/designing-software-cw1/internal/config"
)

func CLI(svc *config.Services) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bankcli",
		Short: "Bank accounting system CLI",
	}
	cmd.AddCommand(
		CreateAccount(svc),
		ApplyIncome(svc),
		ApplyOutcome(svc),
	)

	return cmd
}

func CreateAccount(svc *config.Services) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-account",
		Short: "Create bank account",
	}

	var name string
	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "The name of account")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		acc, err := svc.BankService.CreateAccount(cmd.Context(), name)
		if err != nil {
			return err
		}
		cmd.Printf(`Created an account with ID=%s, Name=%s`, acc.ID, acc.Name)
		return nil
	}
	return cmd
}

func ApplyIncome(svc *config.Services) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply-income",
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

		acc, err := svc.BankService.ApplyOperation(cmd.Context(), services.ApplyOperationRequest{
			AccountID:     accID,
			Amount:        amount,
			OperationType: "income",
		})
		if err != nil {
			return err
		}

		cmd.Printf(`Income operation applied on account with ID=%s, account balance=%d`, acc.ID, acc.Balance)
		return nil
	}
	return cmd
}

func ApplyOutcome(svc *config.Services) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply-outcome",
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

		acc, err := svc.BankService.ApplyOperation(cmd.Context(), services.ApplyOperationRequest{
			AccountID:     accID,
			Amount:        amount,
			OperationType: "outcome",
		})
		if err != nil {
			return err
		}

		cmd.Printf(`Outcome operation applied on account with ID=%s, account balance=%d`, acc.ID, acc.Balance)
		return nil
	}
	return cmd
}

package cli

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/dto"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/services"
)

type BankAccountService interface {
	Get(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error)
	List(ctx context.Context) ([]dto.BankAccountDTO, error)
	CreateAccount(ctx context.Context, name string) (*dto.BankAccountDTO, error)
	Block(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error)
	Unblock(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error)
	Delete(ctx context.Context, id uuid.UUID) (*dto.BankAccountDTO, error)
}

func Account(svc *services.BankAccountService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Operations connected to the bank account",
	}
	cmd.AddCommand(
		getAccount(svc),
		listAccounts(svc),
		createAccount(svc),
		blockAccount(svc),
		unblockAccount(svc),
		deleteAccount(svc),
	)
	return cmd
}

func getAccount(svc BankAccountService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a bank account by its ID",
	}

	var accountIDStr string
	cmd.Flags().StringVarP(&accountIDStr, "id", "i", "", "Account ID")
	cmd.MarkFlagRequired("id")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		accountID, err := uuid.Parse(accountIDStr)
		if err != nil {
			return fmt.Errorf("invalid account ID: %w", err)
		}

		account, err := svc.Get(cmd.Context(), accountID)
		if err != nil {
			return fmt.Errorf("failed to get account: %w", err)
		}

		cmd.Println("Account details:")
		PrettyJSON(cmd, account)
		return nil
	}

	return cmd
}

func listAccounts(svc BankAccountService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all bank accounts",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		accounts, err := svc.List(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list accounts: %w", err)
		}

		cmd.Println("Accounts:")
		PrettyJSON(cmd, accounts)
		return nil
	}

	return cmd
}

func createAccount(svc *services.BankAccountService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create bank account",
	}

	var name string
	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "The name of account")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		acc, err := svc.CreateAccount(cmd.Context(), name)
		if err != nil {
			return err
		}
		cmd.Println(`Created an account:`)
		PrettyJSON(cmd, acc)
		return nil
	}
	return cmd
}

func blockAccount(svc *services.BankAccountService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block",
		Short: "Block bank account",
	}

	var idStr string
	cmd.PersistentFlags().StringVarP(&idStr, "id", "i", "", "The ID of account")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return err
		}

		acc, err := svc.Block(cmd.Context(), id)
		if err != nil {
			return err
		}

		cmd.Println(`Blocked an account:`)
		PrettyJSON(cmd, acc)
		return nil
	}
	return cmd
}

func unblockAccount(svc *services.BankAccountService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unblock",
		Short: "Unblock bank account",
	}

	var idStr string
	cmd.PersistentFlags().StringVarP(&idStr, "id", "i", "", "The ID of account")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return err
		}

		acc, err := svc.Unblock(cmd.Context(), id)
		if err != nil {
			return err
		}

		cmd.Println(`Unblocked an account:`)
		PrettyJSON(cmd, acc)
		return nil
	}
	return cmd
}

func deleteAccount(svc *services.BankAccountService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete bank account",
	}

	var idStr string
	cmd.PersistentFlags().StringVarP(&idStr, "id", "i", "", "The ID of account")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return err
		}

		acc, err := svc.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}

		cmd.Println(`Deleted an account:`)
		PrettyJSON(cmd, acc)
		return nil
	}
	return cmd
}

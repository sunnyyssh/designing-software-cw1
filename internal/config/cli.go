package config

import (
	"github.com/spf13/cobra"
	"github.com/sunnyyssh/designing-software-cw1/internal/cli"
)

func CLI(svc *Services) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bankcli",
		Short: "Bank accounting system CLI",
	}
	cmd.AddCommand(
		cli.Account(svc.BankAccountService),
		cli.Operation(svc.OperationService),
		cli.Category(svc.CategoryService),
	)

	return cmd
}

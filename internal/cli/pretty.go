package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

func PrettyJSON(cmd *cobra.Command, t any) {
	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("\t", "\t")
	if err := encoder.Encode(t); err != nil {
		panic(err)
	}
	cmd.Println()
}

package cmd

import (
	"github.com/spf13/cobra"

	"go-rendang-cli/internal/rendang"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current session stock details",
	Long:  "Show the configured event name and the available starting stock for this run.",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := rendang.NewStore(eventName, totalStock)
		printStatus(cmd.OutOrStdout(), store.Summary())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

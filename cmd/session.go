package cmd

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"go-rendang-cli/internal/rendang"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Run the interactive ordering flow",
	Long:  "Start an interactive rendang ordering session that keeps taking orders until you stop or stock runs out.",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := rendang.NewStore(eventName, totalStock)
		reader := bufio.NewReader(cmd.InOrStdin())
		out := cmd.OutOrStdout()

		printBanner(out, store.Summary())

		for {
			order, err := collectOrder(reader, out, rendang.Order{}, true)
			if err != nil {
				return err
			}

			summary, err := store.PlaceOrder(order)
			if err != nil {
				return err
			}

			fmt.Fprintln(out)
			printOrderReceipt(out, order, summary)

			if summary.RemainingStock == 0 {
				fmt.Fprintln(out, style(out, "Stock is sold out for this session.", ansiBold, ansiRed))
				printSessionSummary(out, summary)
				return nil
			}

			again, err := promptYesNo(reader, out, "Place another order?")
			if err != nil {
				return err
			}
			if !again {
				printSessionSummary(out, summary)
				return nil
			}

			fmt.Fprintln(out)
		}
	},
}

func init() {
	rootCmd.AddCommand(sessionCmd)
}

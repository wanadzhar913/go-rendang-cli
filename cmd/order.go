package cmd

import (
	"bufio"

	"github.com/spf13/cobra"

	"go-rendang-cli/internal/rendang"
)

var (
	firstName string
	lastName  string
	email     string
	packs     uint
	noInput   bool
)

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Place a single rendang order",
	Long:  "Place one order with flags, interactive prompts, or a mix of both.",
	Example: "  rendang order --first-name Faiq --last-name Adzlan --email faiq@example.com --packs 2\n" +
		"  rendang order",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := rendang.NewStore(eventName, totalStock)
		reader := bufio.NewReader(cmd.InOrStdin())
		out := cmd.OutOrStdout()

		order, err := collectOrder(reader, out, rendang.Order{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Packs:     packs,
		}, !noInput)
		if err != nil {
			return err
		}

		summary, err := store.PlaceOrder(order)
		if err != nil {
			return err
		}

		printBanner(out, summary)
		printOrderReceipt(out, order, summary)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(orderCmd)

	orderCmd.Flags().StringVar(&firstName, "first-name", "", "Customer first name")
	orderCmd.Flags().StringVar(&lastName, "last-name", "", "Customer last name")
	orderCmd.Flags().StringVar(&email, "email", "", "Customer email address")
	orderCmd.Flags().UintVar(&packs, "packs", 0, "Number of rendang packs to order")
	orderCmd.Flags().BoolVar(&noInput, "no-input", false, "Disable interactive prompts and require flags")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	eventName  string
	totalStock uint
)

var rootCmd = &cobra.Command{
	Use:   "rendang",
	Short: "A cleaner CLI for festive rendang orders",
	Long: "A Cobra-powered CLI for checking stock, placing rendang orders,\n" +
		"and running an interactive ordering session with better help output.",
	Example: "  rendang status\n" +
		"  rendang order --first-name Faiq --last-name Adzlan --email faiq@example.com --packs 2\n" +
		"  rendang session",
	SilenceUsage: true,
	Version:      "0.1.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.EnableCommandSorting = false

	rootCmd.PersistentFlags().StringVar(&eventName, "event", "Rendang Factory", "Event or storefront name")
	rootCmd.PersistentFlags().UintVar(&totalStock, "stock", 100, "Starting number of rendang packs")

	rootCmd.SetHelpTemplate(`{{with (or .Long .Short)}}{{.}}{{end}}

Usage:
  {{.UseLine}}
{{if .HasAvailableSubCommands}}
Commands:
{{range .Commands}}{{if (and .IsAvailableCommand (not .IsAdditionalHelpTopicCommand))}}  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .Example}}
Examples:
{{.Example}}{{end}}
`)

	rootCmd.InitDefaultCompletionCmd()
}

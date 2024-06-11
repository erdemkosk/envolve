package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type Command interface {
	Execute(cmd *cobra.Command, args []string)
}

var rootCmd = &cobra.Command{
	Use:   "envolve",
	Short: "Envolve CLI is a tool for effortless .env file management.",
	Long: `Envolve is your solution for effortless .env file management.
With Envolve, you can seamlessly gather, arrange, and fine-tune environment variables
across all your projects, ensuring that your configuration data is always at your fingertips
without the risk of loss.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}

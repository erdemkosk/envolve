package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Backs up your current project's .env file",
	Long:  `Backs up your current project's .env file, restores the variables from a global .env file, and creates a symbolic link to the latest environment settings.`,
	Run:   command.CommandFactory(config.SYNC).Execute,
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

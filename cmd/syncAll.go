package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var syncAllCmd = &cobra.Command{
	Use:   "sync-all",
	Short: "Synchronizes all environment variables across projects inside of the folder",
	Long: `Synchronizes all environment variables across projects.
Backs up .env files for all projects, restores variables from a global .env file,
and creates symbolic links to the latest environment settings.`,
	Run: command.CommandFactory(config.SYNCALL).Execute,
}

func init() {
	rootCmd.AddCommand(syncAllCmd)
}

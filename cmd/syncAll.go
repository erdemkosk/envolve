package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var syncAllPath string

var syncAllCmd = &cobra.Command{
	Use:   "sync-all",
	Short: "Synchronizes all environment variables across projects inside of the folder",
	Long: `Synchronizes all environment variables across projects.
Backs up .env files for all projects, restores variables from a global .env file,
and creates symbolic links to the latest environment settings. You can specify a path to sync files from using the --path flag if not it will use current path.`,
	Run: func(cmd *cobra.Command, args []string) {
		command.CommandFactory(config.SYNCALL, syncAllPath).Execute(cmd, args)
	},
}

func init() {
	syncAllCmd.Flags().StringVarP(&syncAllPath, "path", "p", "", "Specify the path of your project to sync env from")
	rootCmd.AddCommand(syncAllCmd)
}

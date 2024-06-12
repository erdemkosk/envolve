package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var syncPath string

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Backs up your current project's .env file",
	Long: `Backs up your current project's .env file, restores the variables from a global .env file, and creates a symbolic link to the latest environment settings.
You can specify a path to sync files from using the --path flag if not it will use current path.`,
	Run: func(cmd *cobra.Command, args []string) {
		command.CommandFactory(config.SYNC, syncPath).Execute(cmd, args)
	},
}

func init() {
	syncCmd.Flags().StringVarP(&syncPath, "path", "p", "", "Specify the path of your project to sync env from")
	rootCmd.AddCommand(syncCmd)
}

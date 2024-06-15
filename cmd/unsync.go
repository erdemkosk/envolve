package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var unsyncCdm = &cobra.Command{
	Use: "unsync",
	Short: `The "unsync" command reverses the synchronization process of environment variables between local and remote projects. It restores the local .env files from the remote location 
		and removes the synchronization link. Additionally, it deletes the project folder from the remote location if it exists.`,
	Long: `The "unsync" command reverses the synchronization process of environment variables between local and remote projects. It restores the local .env files from the remote location 
  		and removes the synchronization link. Additionally, it deletes the project folder from the remote location if it exists.`,
	Run: func(cmd *cobra.Command, args []string) {
		command.CommandFactory(config.UNSYNC, "").Execute(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(unsyncCdm)
}

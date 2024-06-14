package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var editCdm = &cobra.Command{
	Use:   "edit",
	Short: "edit environment variables across projects",
	Long:  `This command provides multiple editing opportunities within the given env keys and values.`,
	Run: func(cmd *cobra.Command, args []string) {
		command.CommandFactory(config.EDIT, syncAllPath).Execute(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(editCdm)
}

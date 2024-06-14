package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var editCdm = &cobra.Command{
	Use:   "edit",
	Short: "Edit environment variables across projects. Using this, you can find all env files according to key and value and change them automatically with a single action.",
	Long:  `This command provides multiple editing opportunities within the given env keys and values.`,
	Run: func(cmd *cobra.Command, args []string) {
		command.CommandFactory(config.EDIT, "").Execute(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(editCdm)
}

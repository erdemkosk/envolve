package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Edit or inspect environment variables in the specified path",
	Long: `Allows editing or inspecting environment variables (.env files) in the specified path.
Lists all the environment variable files (.env) and directories in the specified path and allows
navigation into directories. When an environment variable file is selected, u can edit with internal editor.`,
	Run: command.CommandFactory(config.SHOW).Execute,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

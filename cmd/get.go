package cmd

import (
	config "github.com/erdemkosk/envolve-go/internal"
	command "github.com/erdemkosk/envolve-go/internal/command"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Edit or inspect environment variables in the specified path",
	Long: `Allows editing or inspecting environment variables (.env files) in the specified path.
Lists all the environment variable files (.env) and directories in the specified path and allows
navigation into directories. When an environment variable file is selected, it opens the file with
the default or specified editor for editing.`,
	Run: command.CommandFactory(config.GET).Execute,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

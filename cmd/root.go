package cmd

import (
	"fmt"
	"os"
	"runtime"

	config "github.com/erdemkosk/envolve-go/internal"
	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

type Command interface {
	Execute(cmd *cobra.Command, args []string)
}

var rootCmd = &cobra.Command{
	Use:     "envolve",
	Version: "1.0.17",
	Short:   "Envolve CLI is a tool for effortless .env file management.",
	Long: fmt.Sprintf(`%sEnvolve%s is your solution for effortless .env file management. With %sEnvolve%s, you can seamlessly gather, arrange, and fine-tune environment variables
	across all your projects, ensuring that your configuration data is always at your fingertips without the risk of loss. `,
		config.PASTEL_ORANGE, config.RESET, config.PASTEL_ORANGE, config.RESET),
	// Use colorable for cross-platform ANSI color support
}

var customHelpTemplate = fmt.Sprintf(`{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}

 Usage:
 %s{{.UseLine}}{{if .HasAvailableSubCommands}}  %s

 %sWarning:
 If you have not synchronized any of your projects, everything will be shown as an empty folder in show and edit commands. Edit and show work in already synced projects. %s

%s Info:
 If you want to learn different ways to use a command, for example, just type envolve sync -h for sync. 
 In the description, you will find example information that you can give a path to a folder you want, not the current folder, with --path. %s
  
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  %s{{rpad .Name .NamePadding }}%s  %s{{.Short}}{{end}}{{end}}{{end}} %s
  
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`, config.PASTEL_ORANGE, config.RESET, config.PASTEL_RED, config.RESET, config.PASTEL_CYAN, config.RESET, config.PASTEL_GRAY, config.RESET, config.PASTEL_BLUE, config.RESET)

func Execute() {
	envolvePath := logic.GetEnvolveHomePath()

	err := logic.CreateFolderIfDoesNotExist(envolvePath)
	if err != nil {
		return
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Enable ANSI colors in Windows terminal if supported
	if runtime.GOOS == "windows" && isatty.IsTerminal(os.Stdout.Fd()) {
		rootCmd.SetOut(colorable.NewColorable(os.Stdout))
	}

	rootCmd.SetHelpTemplate(customHelpTemplate)
}

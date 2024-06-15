package cmd

import (
	"fmt"
	"os"

	config "github.com/erdemkosk/envolve-go/internal"
	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/spf13/cobra"
)

type Command interface {
	Execute(cmd *cobra.Command, args []string)
}

var rootCmd = &cobra.Command{
	Use:     "envolve",
	Version: "1.0.13",
	Short:   "Envolve CLI is a tool for effortless .env file management.",
	Long: `` + config.PASTEL_ORANGE + `Envolve ` + config.RESET + `is your solution for effortless .env file management. With ` + config.PASTEL_ORANGE + `Envolve ` + config.RESET + `,you can seamlessly gather, arrange, and fine-tune environment variables
	across all your projects, ensuring that your configuration data is always at your fingertips without the risk of loss. `,
}

var customHelpTemplate = fmt.Sprintf(`{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}

 Usage:
 ` + config.PASTEL_ORANGE + `{{.UseLine}}{{if .HasAvailableSubCommands}}  ` + config.RESET + `

 ` + config.PASTEL_RED + `Warning:
 If you have not synchronized any of your projects, everything will be shown as an empty folder in show and edit commands. Edit and show work in already synced projects. ` + config.RESET + `

` + config.PASTEL_CYAN + ` Info:
 If you want to learn different ways to use a command, for example, just type envolve sync -h for sync. 
 In the description, you will find example information that you can give a path to a folder you want, not the current folder, with --path.  ` + config.RESET + `
  
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  ` + config.PASTEL_GRAY + `{{rpad .Name .NamePadding }}` + config.RESET + `  ` + config.PASTEL_BLUE + ` {{.Short}}{{end}}{{end}}{{end}} ` + config.RESET + `
  
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`)

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
	rootCmd.SetHelpTemplate(customHelpTemplate)
}

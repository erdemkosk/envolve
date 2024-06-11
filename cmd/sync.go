package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Backs up your current project's .env file",
	Long:  `Backs up your current project's .env file, restores the variables from a global .env file, and creates a symbolic link to the latest environment settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		envolvePath := logic.GetEnvolveHomePath()
		currentPath, currentFolderName := logic.GetCurrentPathAndFolder()
		targetPath := filepath.Join(envolvePath, currentFolderName)
		currentEnvFilePath := filepath.Join(currentPath, "/.env")
		targetEnvFilePath := filepath.Join(targetPath, "/.env")

		logic.CreateFolderIfDoesNotExist(targetPath)
		logic.CopyFile(currentEnvFilePath, targetEnvFilePath)
		logic.DeleteFile(currentEnvFilePath)
		logic.Symlink(targetEnvFilePath, currentEnvFilePath)

		log.Printf("Sync work with success!")

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

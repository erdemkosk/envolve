package command

import (
	"log"
	"os"
	"path/filepath"

	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/spf13/cobra"
)

type SyncCommand struct {
	path string
}

func (command *SyncCommand) Execute(cmd *cobra.Command, args []string) {

	envolvePath := logic.GetEnvolveHomePath()
	currentPath, currentFolderName := logic.GetCurrentPathAndFolder(command.path)
	targetPath := filepath.Join(envolvePath, currentFolderName)
	currentEnvFilePath := filepath.Join(currentPath, "/.env")
	targetEnvFilePath := filepath.Join(targetPath, "/.env")

	if _, err := os.Stat(targetEnvFilePath); err == nil {
		log.Println("Error: .env file already exists in the current directory!")
		return
	}

	err := logic.CreateFolderIfDoesNotExist(targetPath)

	if err != nil {
		return
	}

	copyErr := logic.CopyFile(currentEnvFilePath, targetEnvFilePath)

	if copyErr != nil {
		return
	}

	logic.DeleteFile(currentEnvFilePath)
	logic.Symlink(targetEnvFilePath, currentEnvFilePath)

	log.Println("Sync work with success!")

	os.Exit(0)
}

package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/spf13/cobra"
)

type SyncCommand struct {
}

func (command *SyncCommand) Execute(cmd *cobra.Command, args []string) {
	envolvePath := logic.GetEnvolveHomePath()
	currentPath, currentFolderName := logic.GetCurrentPathAndFolder()
	targetPath := filepath.Join(envolvePath, currentFolderName)
	currentEnvFilePath := filepath.Join(currentPath, "/.env")
	targetEnvFilePath := filepath.Join(targetPath, "/.env")

	logic.CreateFolderIfDoesNotExist(targetPath)
	logic.CopyFile(currentEnvFilePath, targetEnvFilePath)
	logic.DeleteFile(currentEnvFilePath)
	logic.Symlink(targetEnvFilePath, currentEnvFilePath)

	fmt.Println("Sync work with success!")

	os.Exit(0)
}

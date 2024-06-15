package command

import (
	"log"
	"os"
	"path/filepath"

	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/spf13/cobra"
)

type UnsyncCommand struct {
}

func (command *UnsyncCommand) Execute(cmd *cobra.Command, args []string) {
	currentPath, currentFolderName := logic.GetCurrentPathAndFolder("")
	envolvePath := logic.GetEnvolveHomePath()
	targetPath := filepath.Join(envolvePath, currentFolderName)
	currentEnvFilePath := filepath.Join(currentPath, ".env")
	targetEnvFilePath := filepath.Join(targetPath, ".env")

	info, err := os.Lstat(currentEnvFilePath)
	if err != nil {
		log.Printf("Error accessing .env file: %v\n", err)
		return
	}

	if info.Mode()&os.ModeSymlink == 0 {
		log.Println("Error: .env file is not a symlink!")
		return
	}

	symlinkTarget, err := os.Readlink(currentEnvFilePath)
	if err != nil {
		log.Printf("Error reading symlink: %v\n", err)
		return
	}

	if symlinkTarget != targetEnvFilePath {
		log.Printf("Error: symlink does not point to the expected target path %s!\n", targetEnvFilePath)
		return
	}

	tempEnvFilePath := currentEnvFilePath + ".tmp"
	copyErr := logic.CopyFile(targetEnvFilePath, tempEnvFilePath)
	if copyErr != nil {
		log.Printf("Error copying .env file back: %v\n", copyErr)
		return
	}

	logic.DeleteFile(currentEnvFilePath)

	renameErr := os.Rename(tempEnvFilePath, currentEnvFilePath)
	if renameErr != nil {
		log.Printf("Error renaming restored .env file: %v\n", renameErr)
		return
	}

	log.Println("Restore operation completed successfully!")

	projectPathToDelete := filepath.Join(envolvePath, currentFolderName)
	if _, err := os.Stat(projectPathToDelete); err == nil {
		log.Printf("Deleting project folder: %s\n", projectPathToDelete)
		err := os.RemoveAll(projectPathToDelete)
		if err != nil {
			log.Printf("Error deleting project folder: %v\n", err)
			return
		}
		log.Printf("Project folder deleted: %s\n", projectPathToDelete)
	} else {
		log.Printf("Project folder %s does not exist in Envolve path.\n", currentFolderName)
	}

	os.Exit(0)
}

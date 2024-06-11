package command

import (
	"log"
	"os"
	"path/filepath"

	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/spf13/cobra"
)

type SyncAllCommand struct{}

func (command *SyncAllCommand) Execute(cmd *cobra.Command, args []string) {
	currentPath, _ := logic.GetCurrentPathAndFolder("")
	envolvePath := logic.GetEnvolveHomePath()

	err := filepath.Walk(currentPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if info.IsDir() && path != currentPath { // Check if it's a directory (excluding current directory)
			log.Printf("Syncing .env file in directory: %s\n", path)

			currentPath, currentFolderName := logic.GetCurrentPathAndFolder(path)
			targetPath := filepath.Join(envolvePath, currentFolderName)
			currentEnvFilePath := filepath.Join(currentPath, "/.env")
			targetEnvFilePath := filepath.Join(targetPath, "/.env")

			if _, err := os.Stat(targetEnvFilePath); err == nil {
				log.Println("Error: .env file already exists in the current directory!")
				return nil
			}

			err := logic.CreateFolderIfDoesNotExist(targetPath)

			if err != nil {
				return nil
			}

			copyErr := logic.CopyFile(currentEnvFilePath, targetEnvFilePath)

			if copyErr != nil {
				return nil
			}

			logic.DeleteFile(currentEnvFilePath)
			logic.Symlink(targetEnvFilePath, currentEnvFilePath)

			log.Printf("Sync completed successfully for directory: %s\n", path)
		}

		return nil
	})

	if err != nil {
		log.Printf("Error walking through directory: %v\n", err)
	}

	os.Exit(0)
}

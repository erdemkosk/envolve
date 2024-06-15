package command

import (
	"log"
	"os"
	"path/filepath"

	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/spf13/cobra"
)

type SyncAllCommand struct {
	path string
}

func (command *SyncAllCommand) Execute(cmd *cobra.Command, args []string) {
	currentPath, _ := logic.GetCurrentPathAndFolder(command.path)
	envolvePath := logic.GetEnvolveHomePath()

	entries, err := os.ReadDir(currentPath)
	if err != nil {
		log.Printf("Error reading directory %q: %v\n", currentPath, err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDirPath := filepath.Join(currentPath, entry.Name())
			envFilePath := filepath.Join(subDirPath, ".env")

			// Check if .env file exists in the subdirectory
			if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
				log.Printf(".env file does not exist in directory: %s, skipping...\n", subDirPath)
				continue
			}

			log.Printf("Syncing .env file in directory: %s\n", subDirPath)

			_, currentFolderName := logic.GetCurrentPathAndFolder(subDirPath)
			targetPath := filepath.Join(envolvePath, currentFolderName)
			targetEnvFilePath := filepath.Join(targetPath, ".env")

			if _, err := os.Stat(targetEnvFilePath); err == nil {
				log.Println("Error: .env file already exists in the target directory!")
				continue
			}

			err := logic.CreateFolderIfDoesNotExist(targetPath)
			if err != nil {
				log.Printf("Error creating target directory: %v\n", err)
				continue
			}

			copyErr := logic.CopyFile(envFilePath, targetEnvFilePath)
			if copyErr != nil {
				log.Printf("Error copying .env file: %v\n", copyErr)
				continue
			}

			logic.DeleteFile(envFilePath)

			logic.Symlink(targetEnvFilePath, envFilePath)

			log.Printf("Sync completed successfully for directory: %s\n", subDirPath)
		}
	}

	os.Exit(0)
}

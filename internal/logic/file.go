package logic

import (
	"io"
	"log"
	"os"
	"path/filepath"

	config "github.com/erdemkosk/envolve-go/internal"
)

func getHomePath() string {
	home, _ := os.UserHomeDir()
	return home
}

func GetEnvolveHomePath() string {
	return filepath.Join(getHomePath(), config.HOME_FOLDER)
}

func contains(names []string, name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}

func ReadDir(path string, excludeNames []string) ([]os.FileInfo, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filteredFiles []os.DirEntry
	for _, file := range files {
		if !contains(excludeNames, file.Name()) {
			filteredFiles = append(filteredFiles, file)
		}
	}

	var fileInfos []os.FileInfo
	for _, file := range filteredFiles {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		fileInfos = append(fileInfos, info)
	}

	return fileInfos, nil
}

func GetCurrentPathAndFolder(optionalPath string) (string, string) {
	path := optionalPath
	if path == "" {
		path, _ = os.Getwd()
	}
	return path, filepath.Base(path)
}

func GetFoldername(path string) string {
	return filepath.Base(filepath.Dir(path))
}

func CreateFolderIfDoesNotExist(homePath string) error {
	if _, err := os.Stat(homePath); os.IsNotExist(err) {
		if err := os.MkdirAll(homePath, 0755); err != nil {
			log.Println("Create folder problem:", err)
			return err
		}
	} else if err != nil {
		log.Println("Error checking directory:", err)
		return err
	}
	return nil
}

func Symlink(source, target string) {
	if err := os.Symlink(source, target); err != nil {
		log.Println("There is a problem with symlink:", err)
	}
}

func CopyFile(sourceFilePath, targetFilePath string) error {
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		log.Println("Source file problem", err)
		return err
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		log.Println("Target file problem", err)
		return err
	}
	defer targetFile.Close()

	if _, err = sourceFile.Seek(0, 0); err != nil {
		log.Println("Seek error", err)
		return err
	}

	if _, err = io.Copy(targetFile, sourceFile); err != nil {
		log.Println("File cannot be copied:", err)
		return err
	}

	return nil
}

func DeleteFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		log.Println("Remove problem", err)
	}
}

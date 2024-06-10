package logic

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	config "github.com/erdemkosk/envolve-go/internal"
)

func getHomePath() string {
	home, _ := os.UserHomeDir()
	return home
}

func GetEnvolveHomePath() string {
	homePath := getHomePath()
	envolvePath := filepath.Join(homePath, config.HOME_FOLDER)

	return envolvePath
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
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filteredFiles []os.FileInfo
	for _, file := range files {
		if !contains(excludeNames, file.Name()) {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles, nil
}

func GetCurrentPathAndFolder() (string, string) {
	path, _ := os.Getwd()
	folder := filepath.Base(path)
	return path, folder
}

func GetFoldername(path string) string {
	folder := filepath.Base(filepath.Dir(path))
	return folder
}

func CreateFolderIfDoesNotExist(homePath string) {
	_, err := os.Stat(homePath)

	if os.IsNotExist(err) {
		err := os.Mkdir(homePath, 0755)

		if err != nil {
			fmt.Println("Create folder problem:", err)
			return
		}
	}
}

func Symlink(source string, target string) {
	err := os.Symlink(source, target)

	if err != nil {
		fmt.Println("There is a problem with symlink:", err)
		return
	}
}

func CopyFile(sourceFilePath string, targetFilePath string) {
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("Source file problem", err)
		return
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		fmt.Println("Target file problem", err)
		return
	}
	defer targetFile.Close()

	_, err = sourceFile.Seek(0, 0)
	if err != nil {
		fmt.Println("Seek error", err)
		return
	}

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		fmt.Println("File cannot copied:", err)
		return
	}
}

func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Remove problem", err)
		return
	}

}

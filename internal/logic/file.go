package logic

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	config "github.com/erdemkosk/envolve-go/internal"
	"github.com/rivo/tview"
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

func CreateFolderIfDoesNotExist(homePath string) error {
	_, err := os.Stat(homePath)
	if os.IsNotExist(err) {
		er := os.MkdirAll(homePath, 0755)
		if er != nil {
			log.Println("Create folder problem:", err)
			return err
		}
	} else if err != nil {
		log.Println("Error checking directory:", err)
		return err
	}

	return nil
}

func Symlink(source string, target string) {
	err := os.Symlink(source, target)

	if err != nil {
		log.Println("There is a problem with symlink:", err)
		return
	}
}

func CopyFile(sourceFilePath string, targetFilePath string) error {
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

	_, err = sourceFile.Seek(0, 0)
	if err != nil {
		log.Println("Seek error", err)
		return err
	}

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		log.Println("File cannot copied:", err)
		return err
	}

	return nil
}

func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("Remove problem", err)
		return
	}

}

func ShowFileContent(file string, rightBox *tview.TextArea) {
	content, err := os.ReadFile(file)
	if err != nil {
		rightBox.SetText("Error reading file: "+err.Error(), true)
		return
	}

	rightBox.SetText(string(content), true)
}

func SaveFileContent(file string, rightBox *tview.TextArea) {
	text := rightBox.GetText()
	if err := os.WriteFile(file, []byte(text), 0644); err != nil {
		rightBox.SetText("Error saving file: "+err.Error(), true)
	}
}

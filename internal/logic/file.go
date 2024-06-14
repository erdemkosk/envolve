package logic

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

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

func ReadEnvFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envs := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := parts[1]
		envs[key] = value
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return envs, nil
}

type EnvFile struct {
	Path    string
	Folder  string
	EnvVars map[string]string
}

func CollectEnvFiles(envDir string) ([]EnvFile, error) {
	var envFiles []EnvFile
	err := filepath.Walk(envDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".env") {
			envVars, err := ReadEnvFile(path)
			if err != nil {
				return err
			}
			envFiles = append(envFiles, EnvFile{
				Path:    path,
				Folder:  filepath.Base(filepath.Dir(path)),
				EnvVars: envVars,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return envFiles, nil
}

func UpdateEnvFiles(envFiles []EnvFile, key, newValue string) error {
	for i := range envFiles {
		envFile := &envFiles[i]
		if _, exists := envFile.EnvVars[key]; exists {
			envFile.EnvVars[key] = newValue
			if err := WriteEnvFile(envFile.Path, envFile.EnvVars); err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateEnvFilesWithValue(envFiles []EnvFile, value, newValue string) error {
	for _, envFile := range envFiles {
		for key, envValue := range envFile.EnvVars {
			if envValue == value {
				envFile.EnvVars[key] = newValue
				if err := WriteEnvFile(envFile.Path, envFile.EnvVars); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func WriteEnvFile(path string, envVars map[string]string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range envVars {
		_, err := fmt.Fprintf(writer, "%s=%s\n", key, value)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

type EditCommand struct{}

func (command EditCommand) Execute(cmd *cobra.Command, args []string) {
	envDir := logic.GetEnvolveHomePath()

	envFiles, err := logic.CollectEnvFiles(envDir)
	if err != nil {
		log.Fatalf("Failed to collect env files: %v", err)
	}

	app := tview.NewApplication()

	resultBox := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	resultBox.SetTitle("Matches").SetBorder(true)

	searchBox := tview.NewInputField().
		SetLabel("Search Key: ").
		SetFieldWidth(20).
		SetChangedFunc(func(text string) {
			if text == "" {
				return
			}
			updateResults(resultBox, envFiles, text)
		}).
		SetAutocompleteFunc(func(currentText string) (entries []string) {
			if currentText == "" || len(envFiles) == 0 {
				return nil
			}
			// Get all keys that start with the current text
			for key := range envFiles[0].EnvVars { // Assuming envFiles is not empty
				if strings.HasPrefix(key, currentText) {
					entries = append(entries, key)
				}
			}
			return entries
		})

	valueSearchBox := tview.NewInputField().
		SetLabel("Search Value: ").
		SetFieldWidth(20).
		SetChangedFunc(func(text string) {
			if text == "" {
				return
			}
			updateResultsByValue(resultBox, envFiles, text)
		}).
		SetAutocompleteFunc(func(currentText string) (entries []string) {
			if currentText == "" || len(envFiles) == 0 {
				return nil
			}

			seenValues := make(map[string]bool)
			for _, envFile := range envFiles {
				for _, value := range envFile.EnvVars {
					if strings.HasPrefix(value, currentText) && !seenValues[value] {
						seenValues[value] = true
						entries = append(entries, value)
					}
				}
			}
			return entries
		})

	valueBox := tview.NewInputField().
		SetLabel("New Value: ").
		SetFieldWidth(20)

	form := tview.NewForm().
		AddFormItem(searchBox).
		AddFormItem(valueSearchBox).
		AddFormItem(valueBox).
		AddButton("Update", func() {
			// This is the action when the "Update" button is pressed
			key := searchBox.GetText()
			value := valueSearchBox.GetText()
			newValue := valueBox.GetText()

			if key == "" && value == "" {
				resultBox.SetText("[red]Key or search value must be provided")
				return
			}

			var matchingFiles []logic.EnvFile
			if key != "" {
				matchingFiles = filterEnvFilesByKey(envFiles, key)
			} else if value != "" {
				matchingFiles = filterEnvFilesByValue(envFiles, value)
			}

			if len(matchingFiles) == 0 {
				if key != "" {
					resultBox.SetText("[yellow]No matching keys found")
				} else {
					resultBox.SetText("[yellow]No matching values found")
				}
				return
			}

			err := logic.UpdateEnvFiles(matchingFiles, key, newValue)
			if err != nil {
				resultBox.SetText("[red]Failed to update env files: " + err.Error())
				return
			}

			resultBox.SetText("[green]Successfully updated env files")
		}).
		AddButton("Update by Value", func() { // Add a new button for updating by value
			value := valueSearchBox.GetText()
			newValue := valueBox.GetText()

			if value == "" {
				resultBox.SetText("[red]Search value must be provided")
				return
			}

			matchingFiles := filterEnvFilesByValue(envFiles, value)

			err := logic.UpdateEnvFilesWithValue(matchingFiles, value, newValue)
			if err != nil {
				resultBox.SetText("[red]Failed to update env files by value: " + err.Error())
				return
			}

			resultBox.SetText("[green]Successfully updated env files by value")
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	// Add empty TextView for spacing

	form.SetTitle("Search").SetBorder(true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(form, 0, 1, false).
		AddItem(resultBox, 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func filterEnvFilesByKey(envFiles []logic.EnvFile, key string) []logic.EnvFile {
	if len(envFiles) == 0 {
		return nil
	}
	var matchingFiles []logic.EnvFile
	for _, envFile := range envFiles {
		if _, exists := envFile.EnvVars[key]; exists {
			matchingFiles = append(matchingFiles, envFile)
		}
	}
	return matchingFiles
}

func filterEnvFilesByValue(envFiles []logic.EnvFile, value string) []logic.EnvFile {
	if len(envFiles) == 0 {
		return nil
	}

	var matchingFiles []logic.EnvFile
	for _, envFile := range envFiles {
		for _, envValue := range envFile.EnvVars {
			if envValue == value {
				matchingFiles = append(matchingFiles, envFile)
				break
			}
		}
	}
	return matchingFiles
}

func updateResults(resultBox *tview.TextView, envFiles []logic.EnvFile, searchKey string) {
	resultBox.Clear()
	if searchKey == "" || len(envFiles) == 0 {
		resultBox.SetText("[yellow]No matching keys found")
		return
	}

	var results []string
	for _, envFile := range envFiles {
		for key, value := range envFile.EnvVars {
			if key == searchKey {
				results = append(results, fmt.Sprintf("%s (%s): %s=%s", envFile.Folder, envFile.Path, key, value))
			}
		}
	}

	if len(results) == 0 {
		resultBox.SetText("[yellow]No matching keys found")
	} else {
		resultBox.SetText("[green]" + strings.Join(results, "\n"))
	}
}

func updateResultsByValue(resultBox *tview.TextView, envFiles []logic.EnvFile, searchValue string) {
	resultBox.Clear()
	if searchValue == "" || len(envFiles) == 0 {
		resultBox.SetText("[yellow]No matching values found")
		return
	}

	var results []string
	for _, envFile := range envFiles {
		for key, value := range envFile.EnvVars {
			if value == searchValue {
				results = append(results, fmt.Sprintf("%s (%s): %s=%s", envFile.Folder, envFile.Path, key, value))
			}
		}
	}

	if len(results) == 0 {
		resultBox.SetText("[yellow]No matching values found")
	} else {
		resultBox.SetText("[green]" + strings.Join(results, "\n"))
	}
}

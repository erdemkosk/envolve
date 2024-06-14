package handler

import (
	"os"

	"github.com/rivo/tview"
)

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

package handler

import (
	"log"
	"os"
	"os/exec"
)

func OpenWithEditorCommand(selectedPath string) *exec.Cmd {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	cmd := exec.Command(editor, selectedPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func Exec(cmd *exec.Cmd) {
	err := cmd.Run()

	if err != nil {
		log.Printf("Error running command: %v", err)
		return
	}
}

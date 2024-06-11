package command

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	config "github.com/erdemkosk/envolve-go/internal"
	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

type ShowCommand struct {
}

func (command ShowCommand) Execute(cmd *cobra.Command, args []string) {
	var mu sync.Mutex
	var currentFilePath string
	var modal *tview.Modal

	envolvePath := logic.GetEnvolveHomePath()

	app := tview.NewApplication()

	tree := tview.NewTreeView().
		SetRoot(tview.NewTreeNode(envolvePath).SetColor(config.MAIN_COLOR)).
		SetCurrentNode(tview.NewTreeNode(envolvePath).SetColor(config.MAIN_COLOR))

	tree.SetTitle("Envs").SetBorder(true)

	rightBox := tview.NewTextArea().
		SetPlaceholder("Enter text here...")
	rightBox.SetTitle("Editor").SetBorder(true)

	help1 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Info(U can use mouse) [yellow]ctrl+N[white]:Move Editor [yellow]ctrl+B[white]:Move Envs [yellow]ctrl+S[white]:Save Changes `)

	addNodes := func(target *tview.TreeNode, path string) {
		files, err := logic.ReadDir(path, config.EXCLUDED_FILES)
		if err != nil {
			log.Fatalf("Cannot read directory: %v", err)
		}

		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name()))
			if file.IsDir() {
				node.SetColor(config.FOLDER_COLOR)
				node.SetSelectable(true)
				node.SetExpanded(false)
			} else {
				node.SetColor(config.FILE_COLOR)
				node.SetSelectable(true)
			}
			target.AddChild(node)
		}
	}

	addNodes(tree.GetRoot(), envolvePath)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		mu.Lock()
		defer mu.Unlock()

		reference := node.GetReference()
		if reference != nil {
			selectedPath := reference.(string)
			stat, err := os.Stat(selectedPath)
			if err != nil {
				log.Printf("Error retrieving details of %s: %v", selectedPath, err)
				return
			}
			if stat.IsDir() {
				if node.IsExpanded() {
					node.Collapse()
				} else {
					node.ClearChildren()
					addNodes(node, selectedPath)
					node.Expand()
				}
			} else {
				currentFilePath = selectedPath
				logic.ShowFileContent(selectedPath, rightBox)
			}
		}
	})
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexColumn).
			AddItem(tree, 0, 1, false).
			AddItem(rightBox, 0, 2, false), 0, 3, false).
		AddItem(help1, 1, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyCtrlN:
			app.SetFocus(rightBox)
		case tcell.KeyCtrlB:
			app.SetFocus(tree)
		case tcell.KeyCtrlS:
			if modal == nil {
				modal = tview.NewModal().
					SetText("Do you want to save changes?").
					AddButtons([]string{"Yes", "Cancel"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Yes" {
							if currentFilePath != "" {
								logic.SaveFileContent(currentFilePath, rightBox)
							}
							app.Stop()
						}

					})

				flex.AddItem(modal, 1, 0, false)
				app.SetFocus(modal)
			} else {
				app.SetFocus(modal)
			}
		}
		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(tree).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

package command

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	config "github.com/erdemkosk/envolve-go/internal"
	"github.com/erdemkosk/envolve-go/internal/handler"
	"github.com/erdemkosk/envolve-go/internal/logic"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

type GetCommand struct {
}

func (command *GetCommand) Execute(cmd *cobra.Command, args []string) {
	var mu sync.Mutex

	envolvePath := logic.GetEnvolveHomePath()

	app := tview.NewApplication()

	tree := tview.NewTreeView().
		SetRoot(tview.NewTreeNode(envolvePath).SetColor(config.MAIN_COLOR)).
		SetCurrentNode(tview.NewTreeNode(envolvePath).SetColor(config.MAIN_COLOR))

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

				cmd := handler.OpenWithEditorCommand(selectedPath)

				app.Stop()

				handler.Exec(cmd)
			}
		}
	})

	if err := app.SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}

	os.Exit(0)
}

package cmd

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

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Backs up your current project's .env file",
	Long:  `Backs up your current project's .env file, restores the variables from a global .env file, and creates a symbolic link to the latest environment settings.`,
	Run: func(cmd *cobra.Command, args []string) {

		var mu sync.Mutex

		envolvePath := logic.GetEnvolveHomePath()

		app := tview.NewApplication()

		tree := tview.NewTreeView().
			SetRoot(tview.NewTreeNode(envolvePath).SetColor(config.MAIN_COLOR)).
			SetCurrentNode(tview.NewTreeNode(envolvePath).SetColor(config.MAIN_COLOR))

		tree.SetBorder(true)

		rightBox := tview.NewTextView().SetTextAlign(tview.AlignLeft).SetScrollable(true)

		rightBox.SetBorder(true)

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
					showFileContent(selectedPath, rightBox)
				}
			}
		})

		flex := tview.NewFlex().
			AddItem(tree, 0, 1, false).
			AddItem(rightBox, 0, 1, false)

		if err := app.SetRoot(flex, true).SetFocus(tree).Run(); err != nil {
			panic(err)
		}

		app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyRight {
				app.SetFocus(rightBox)
			}
			return event
		})
	},
}

func showFileContent(file string, rightBox *tview.TextView) {
	content, err := os.ReadFile(file)
	if err != nil {
		rightBox.SetText("Error reading file: " + err.Error())
		return
	}

	rightBox.SetText(string(content))
}

func init() {
	rootCmd.AddCommand(testCmd)
}

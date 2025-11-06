package ui

import (
	"aaleman/dirsize/dir"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GenerateTree(entry *dir.Entry) *tview.TreeView {
	root := tview.NewTreeNode(entry.String()).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(setSelectedFunc)

	tree.SetInputCapture(
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyRune {
				switch event.Rune() {
				case 'l':
					node := tree.GetCurrentNode()
					if node != nil {
						reference := node.GetReference()
						entry := reference.(*dir.Entry)

						if entry.IsDir {
							if len(entry.Files) != len(node.GetChildren()) {
								add(node, entry)
							}
							node.Expand()
						}
					}
					return nil
				case 'h':
					node := tree.GetCurrentNode()
					if node != nil && node.IsExpanded() {
						node.SetExpanded(false)
					}

					return nil
				}
			}
			return event
		})

	add(root, entry)

	return tree
}

func add(target *tview.TreeNode, entry *dir.Entry) {
	target.SetReference(entry)

	for _, file := range entry.Files {
		node := tview.NewTreeNode(file.String()).
			SetReference(file.Path).
			SetSelectable(true)
			// SetSelectable(file.IsDir)
		if file.IsDir {
			node.SetColor(tcell.ColorGreen)
		}
		node.SetReference(&file)
		target.AddChild(node)
	}
}

func setSelectedFunc(node *tview.TreeNode) {
	reference := node.GetReference()
	if reference == nil {
		return
	}
	children := node.GetChildren()
	if len(children) == 0 {
		// Load and show files in this directory.
		entry := reference.(*dir.Entry)
		add(node, entry)
	} else {
		// Collapse if visible, expand if collapsed.
		node.SetExpanded(!node.IsExpanded())
	}
}

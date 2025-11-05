package ui

import (
	"aaleman/dirsize/dir"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GenerateTree(entry *dir.Entry) *tview.TreeView {
	root := tview.NewTreeNode(entry.String()).SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	add(root, entry)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
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
	})

	return tree
}

func add(target *tview.TreeNode, entry *dir.Entry) {
	target.SetReference(entry)
	for _, file := range entry.Files {
		node := tview.NewTreeNode(file.String()).SetReference(file.Path).SetSelectable(file.IsDir)
		if file.IsDir {
			node.SetColor(tcell.ColorGreen)
		}
		node.SetReference(&file)
		target.AddChild(node)
	}
}

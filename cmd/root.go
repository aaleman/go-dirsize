/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"aaleman/dirsize/dir"
	"aaleman/dirsize/ui"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dirsize path",
	Short: "dirsize",
	Args:  cobra.MaximumNArgs(1),
	// SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultPath := "."
		if len(args) == 1 {
			defaultPath = args[0]
		}

		hidden, err := cmd.Flags().GetBool("hidden")
		if err != nil {
			return err
		}

		searchConfig := dir.SearchConfig{
			Hidden: hidden,
		}

		entries := dir.ReadFolder(defaultPath, searchConfig)
		// entries.PrintRec()

		tree := ui.GenerateTree(entries)

		if err := tview.NewApplication().SetRoot(tree, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("hidden", "H", false, "Include hidden files")
}

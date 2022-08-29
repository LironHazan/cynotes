package cmd

import (
	cynotes "cynotes/pkg"
	promptUiUtils "cynotes/pkg/ui"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit note",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := promptUiUtils.BasicPrompt("Enter a note name")
		if err != nil {
			return
		}
		cynotes.EditNote(name)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

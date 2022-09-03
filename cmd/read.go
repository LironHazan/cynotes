package cmd

import (
	cynotes "cynotes/pkg"
	promptUiUtils "cynotes/pkg/ui"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Choose a revision to read",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		note, err := promptUiUtils.BasicPrompt("Enter the note name:")
		if err != nil {
			return
		}
		cynotes.Read(note)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

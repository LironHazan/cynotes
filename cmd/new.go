package cmd

import (
	promptUiUtils "cynotes/pkg/ui"
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new note",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := promptUiUtils.PromptFileName()
		fmt.Printf("%s \n", filename)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

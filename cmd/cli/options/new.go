package options

import (
	cynotes "cynotes/pkg"
	promptUiUtils "cynotes/pkg/ui"
	"github.com/spf13/cobra"
	"log"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new note",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		note, _ := promptUiUtils.PromptFileName()
		err := cynotes.New(note)
		if err != nil {
			log.Printf("error creating new note")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

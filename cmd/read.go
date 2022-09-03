package cmd

import (
	cynotes "cynotes/pkg"
	promptUiUtils "cynotes/pkg/ui"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

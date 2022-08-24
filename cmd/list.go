package cmd

import (
	cynotes "cynotes/pkg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list revisions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cynotes.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

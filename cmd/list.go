package cmd

import (
	cynotes "cynotes/pkg"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list revisions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		notes := cynotes.List()
		fmt.Println(notes)

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

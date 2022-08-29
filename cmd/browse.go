package cmd

import (
	cynotes "cynotes/pkg"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Open the repo in the browser",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cynotes.Browse()
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}

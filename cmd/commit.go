package cmd

import (
	cynotes "cynotes/pkg"
	promptUiUtils "cynotes/pkg/ui"
	"fmt"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "commit",
	Short: "A safe saving file versions like git commit",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			passwd, _ := promptUiUtils.PromptPasswdInput()
			cynotes.Commit(args[0], passwd)
		} else {
			fmt.Println("Please provide both filepath and a passphrase.")
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().BoolP("filepath", "f", true, "Help message for toggle")
}

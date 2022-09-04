package options

import (
	cynotes "cynotes/pkg"
	fsutils "cynotes/pkg/fs"
	promptUiUtils "cynotes/pkg/ui"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cynotes",
	Short: "Welcome to cynotes",
	Long:  ``,
}

func InitSecretWorkspace() {
	path, _ := fsutils.GetCYNotesPath()
	if !fsutils.IsPathExists(path + "/.init.json") {
		user, err := promptUiUtils.PromptGitClient("Enter git username")
		if err != nil {
			return
		}

		repo, err := promptUiUtils.PromptGitClient("Enter cynotes repo")
		if err != nil {
			return
		}

		err = cynotes.InitCYNotes(user, repo)
		if err != nil {
			return
		}
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	ascii := promptUiUtils.CynotsArt()
	fmt.Printf(ascii)

	InitSecretWorkspace()
}

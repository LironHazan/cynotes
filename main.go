package main

import (
	"cynotes/cmd"
	cynotes "cynotes/pkg"
	promptUiUtils "cynotes/pkg/ui"
	"fmt"
)

func main() {
	ascii := promptUiUtils.CynotsArt()
	fmt.Printf(ascii)

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

	cmd.Execute()
}

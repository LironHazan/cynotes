package main

import (
	"cynotes/cmd"
	cynotes "cynotes/pkg"
	fsutils "cynotes/pkg/fs"
	promptUiUtils "cynotes/pkg/ui"
	"fmt"
	"sync"
)

func main() {
	ascii := promptUiUtils.CynotsArt()
	fmt.Printf(ascii)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
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
		wg.Done()
	}()
	cmd.Execute()
	wg.Wait()
}

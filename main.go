package main

import (
	"cynotes/cmd"
	fsutils "cynotes/pkg/fs"
	promptUiUtils "cynotes/pkg/ui"
	"fmt"
)

func main() {
	err := fsutils.InitCYNotes()
	if err != nil {
		return
	}
	ascii := promptUiUtils.CynotsArt()
	fmt.Printf(ascii)
	cmd.Execute()
}

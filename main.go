package main

import (
	"cynotes/cmd"
	fsutils "cynotes/pkg/fs"
)

func main() {
	err := fsutils.InitCYNotes()
	if err != nil {
		return
	}
	cmd.Execute()
}

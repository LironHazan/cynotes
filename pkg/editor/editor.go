package editor

import (
	"log"
	"os"
	"os/exec"
)

func ViEdit(filepath string) error {
	cmd := exec.Command("vi", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Error while editing. Error: %v\n", err)
		return err
	}
	log.Printf("Successfully edited.")
	return nil

}

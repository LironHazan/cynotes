package git

import (
	"fmt"
	"log"
	"os/exec"
)

// gh repo clone LironHazan/_cynotes
func Clone(user string, repo string, target string) error {
	cmd := exec.Command("gh", "repo", "clone", user+"/"+repo)
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf(string(out))
	return nil
}

func Add(target string, file string) {
	cmd := exec.Command("git", "add", file)
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Add Error: %s", err)
		return
	}
	fmt.Printf(string(out))
}

func Commit(target string) {
	cmd := exec.Command("git", "commit", "-m", "secrets")
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Commit Error: %s", err)
		return
	}
	fmt.Printf(string(out))
}

func Push(target string) {
	cmd := exec.Command("git", "push", "origin", "master")
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Push Error: %s", err)
		return
	}
	fmt.Printf(string(out))
}

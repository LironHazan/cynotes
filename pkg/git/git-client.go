package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func hasExe(exeName string) bool {
	if _, err := exec.Command(exeName).Output(); err != nil {
		log.Printf("Install %s %s", exeName, err)
		return false
	}
	return true
}

// gh repo clone LironHazan/_cynotes
func Clone(user string, repo string, target string) error {
	if !hasExe("gh") {
		os.Exit(1)
	}
	cmd := exec.Command("gh", "repo", "clone", user+"/"+repo)
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf(string(out))
	return nil
}

func Browse(target string) error {
	if !hasExe("gh") {
		os.Exit(1)
	}
	cmd := exec.Command("gh", "browse")
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf(string(out))
	return nil
}

func Add(target string, file string) {
	if !hasExe("git") {
		os.Exit(1)
	}
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
	if !hasExe("git") {
		os.Exit(1)
	}
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
	if !hasExe("git") {
		os.Exit(1)
	}
	cmd := exec.Command("git", "push", "origin", "master")
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Push Error: %s", err)
		return
	}
	fmt.Printf(string(out))
}

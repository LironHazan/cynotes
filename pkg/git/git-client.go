package git

import (
	"fmt"
	"os/exec"
)

// gh repo clone LironHazan/_cynotes
func Clone(user string, repo string, target string) {
	cmd := exec.Command("gh", "repo", "clone", user+"/"+repo)
	cmd.Dir = target
	out, err := cmd.Output()
	if err != nil {
		return
	}
	fmt.Printf(string(out))
}

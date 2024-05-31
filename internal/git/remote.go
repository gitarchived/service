package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func RemoteLastCommit(url string) (string, error) {
	cmd := exec.Command("git", "ls-remote", url)

	// Set environment variables to avoid asking for credentials
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_ASKPASS=%s", "echo"))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_TERMINAL_PROMPT=0"))

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.Fields(string(out))[0], nil
}

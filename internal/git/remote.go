package git

import (
	"os/exec"
	"strings"
)

func RemoteLastCommit(url string) (string, error) {
	cmd := exec.Command("git", "ls-remote", url)

	// Set environment variables to avoid asking for credentials
	cmd.Env = append(cmd.Env, "GIT_ASKPASS=echo")
	cmd.Env = append(cmd.Env, "GIT_TERMINAL_PROMPT=0")

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.Fields(string(out))[0], nil
}

package git

import (
	"fmt"
	"os/exec"
)

func Clone(url string, dir string) error {
	cmd := exec.Command("git", "clone", "--depth=100", url, dir)

	// Set environment variables to avoid asking for credentials
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_ASKPASS=%s", "echo"))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_TERMINAL_PROMPT=0"))

	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}

package git

import (
	"os/exec"
)

func Clone(url string, dir string) error {
	cmd := exec.Command("git", "clone", "--depth=100", url, dir)

	// Set environment variables to avoid asking for credentials
	cmd.Env = append(cmd.Env, "GIT_ASKPASS=echo")
	cmd.Env = append(cmd.Env, "GIT_TERMINAL_PROMPT=0")

	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}

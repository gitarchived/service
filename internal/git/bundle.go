package git

import (
	"fmt"
	"os/exec"
)

func Bundle(name string, id uint) error {
	cmd := exec.Command("git", "bundle", "create", fmt.Sprintf("%d.bundle", id), "HEAD")

	// Set environment variables to avoid asking for credentials
	cmd.Env = append(cmd.Env, "GIT_ASKPASS=echo")
	cmd.Env = append(cmd.Env, "GIT_TERMINAL_PROMPT=0")

	cmd.Dir = fmt.Sprintf("./%s", name)

	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}

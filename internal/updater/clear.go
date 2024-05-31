package updater

import (
	"os"
)

func Clear(path string) error {
	err := os.RemoveAll(path)

	if err != nil {
		return err
	}

	err = os.RemoveAll(string(path[0]))

	if err != nil {
		return err
	}

	return nil
}

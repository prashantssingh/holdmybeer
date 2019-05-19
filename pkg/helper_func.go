package pkg

import (
	"os/exec"

	"golang.org/x/sys/unix"
)

func isRoot() bool {
	return unix.Geteuid() == 0
}

func runCommand(workdir string, subCommand string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", subCommand)
	if workdir != "" {
		cmd.Dir = workdir
	}

	if out, err := cmd.CombinedOutput(); err != nil {
		return out, err
	}

	return nil, nil
}

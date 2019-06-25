package pkg

import (
	"os/exec"
	"strings"

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

func checkVersion(cmdCheckVersion string) (bool, string) {
	strNotFound := "not found"
	out, err := runCommand("", cmdCheckVersion)
	if err != nil {
		return false, ""
	}

	outStr := string(out)
	if strings.Contains(outStr, strNotFound) {
		return false, ""
	}

	return true, outStr
}

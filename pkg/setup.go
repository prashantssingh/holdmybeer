package pkg

import (
	"fmt"
)

const (
	installCmdCurl = `sudo apt install -y curl`
	installCmdVim  = `sudo apt install -y vim`
	installCmdHtop = `sudo apt install -y htop`
)

// SetupBareMinimum installs curl, vim and htop as part of the setup
func SetupBareMinimum() error {

	fmt.Print(" >>>>>  installing curl... ")
	if _, err := runCommand("", installCmdCurl); err != nil {
		return fmt.Errorf("setup: command failed with err: %+v", err)
	}
	fmt.Println("done")

	fmt.Print(" >>>>>  installing htop... ")
	if _, err := runCommand("", installCmdHtop); err != nil {
		return fmt.Errorf("setup: command failed with err: %+v", err)
	}
	fmt.Println("done")

	fmt.Print(" >>>>>  installing vim... ")
	if _, err := runCommand("", installCmdVim); err != nil {
		return fmt.Errorf("setup: command failed with err: %+v", err)
	}
	fmt.Println("done")
	return nil
}

package pkg

import (
	"fmt"
	"strings"
)

const (
	hostOS   = "linux"
	hostArch = "amd64"

	workDir = "~/hmb"
)

var (
	// configurations for installing go
	goVersion = "1.12.5"
	goRoot    = "$HOME/hmb/go"
	goPath    = "$HOME/workspace/go/src"
)

// RunInstaller installs a language or a framework based on the flag-value provided to hmb while executing
// the binary on the terminal.
func RunInstaller(langStr string) error {
	arr := strings.Split(langStr, ":")
	lang := arr[0]

	var version string
	if len(arr) > 0 {
		version = arr[1]
	}

	switch lang {
	case "go":
		if err := installGo(version); err != nil {
			return fmt.Errorf("install: failed to install go, err: %+v", err)
		}
	default:
		return fmt.Errorf("install: unknown or unsupported language or tool. Please refer doc for the list of supported language or tool")
	}
	return nil
}

func installGo(version string) error {
	if version != "" {
		goVersion = version
	}

	// https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
	fmt.Print(" >>>>>  downloading go... ")
	downloadLink := fmt.Sprintf("https://dl.google.com/go/go%s.%s-%s.tar.gz", goVersion, hostOS, hostArch)
	cmd := fmt.Sprintf("curl -sSL %s", downloadLink)
	if _, err := runCommand(workDir, cmd); err != nil {
		return fmt.Errorf("setup: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  extracting download... ")
	cmd = fmt.Sprintf("tar -xf go%s.%s-%s.tar.gz", goVersion, hostOS, hostArch)
	if _, err := runCommand(workDir, cmd); err != nil {
		return fmt.Errorf("setup: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  setting-up directories for workspace... ")
	cmd = "mkdir -p ~/workspace/go/{src,pkg,bin}"
	if _, err := runCommand("", cmd); err != nil {
		return fmt.Errorf("setup: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  configuring go environment... ")
	cmd = fmt.Sprintf("printf \"\n\nexport GOROOT=%s\" >> $HOME/.profile && printf \"\nexport GOPATH=%s\" >> $HOME/.profile && printf \"\nexport PATH=$PATH:$GOROOT/bin:$GOPATH/bin\" >> $HOME/.profile", goRoot, goPath)
	if _, err := runCommand("", cmd); err != nil {
		return fmt.Errorf("setup: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  loading go environment... ")
	_, _ = runCommand("", "source $HOME/.profile")
	fmt.Println(" done")

	fmt.Println("\n Go installation was successful. Run 'go version' to check version installed and run 'go env' to check go-specific environment")
	return nil
}

package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	hostOS   = "linux"
	hostArch = "amd64"

	workDir = "~/hmb"
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
	case "go", "golang":
		if err := installGo(version); err != nil {
			return fmt.Errorf("install: failed to install go, err: %+v", err)
		}
	case "node", "nodejs":
		if err := installGo(version); err != nil {
			return fmt.Errorf("install: failed to install go, err: %+v", err)
		}
	default:
		return fmt.Errorf("install: unknown or unsupported language or tool. Please refer doc for the list of supported language or tool")
	}
	return nil
}

func installGo(version string) error {
	// configurations for installing go
	var (
		goVersion = "1.12.5"
		goRoot    = "$HOME/hmb/go"
		goPath    = "$HOME/workspace/go/src"
	)

	if version != "" {
		goVersion = version
	}

	// Check if go is already installed.If yes, check version and only install newer version of Go
	// after prompting the user for his permission. If installed Go version is already the newest
	// then notify user and exit.
	if isInstalled, version := checkVersion("go version"); isInstalled {
		fmt.Printf("go is already installed. Found - %s\n", version)

	label:
		fmt.Printf("Continue? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("install: input read failed with err: %+v", err)
		}

		// handle user-input
		switch input {
		case "":
			fmt.Print("\nPlease enter 'n' to abort installation and 'y' to continue..")
			goto label

		case "n", "N":
			return nil

		case "y", "Y":
			break

		default:
			fmt.Printf("\ninvalid input: %s. Please enter 'n' to abort installation and 'y' to continue..", input)
		}
	}

	// https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
	fmt.Print(" >>>>>  downloading go... ")
	downloadLink := fmt.Sprintf("https://dl.google.com/go/go%s.%s-%s.tar.gz", goVersion, hostOS, hostArch)
	cmd := fmt.Sprintf("curl -sSL %s", downloadLink)
	if _, err := runCommand(workDir, cmd); err != nil {
		return fmt.Errorf("install: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  extracting download... ")
	cmd = fmt.Sprintf("tar -xf go%s.%s-%s.tar.gz", goVersion, hostOS, hostArch)
	if _, err := runCommand(workDir, cmd); err != nil {
		return fmt.Errorf("install: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  setting-up directories for workspace... ")
	cmd = "mkdir -p ~/workspace/go/{src,pkg,bin}"
	if _, err := runCommand("", cmd); err != nil {
		return fmt.Errorf("install: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  configuring go environment... ")
	cmd = fmt.Sprintf("printf \"\n\nexport GOROOT=%s\" >> $HOME/.profile && printf \"\nexport GOPATH=%s\" >> $HOME/.profile && printf \"\nexport PATH=$PATH:$GOROOT/bin:$GOPATH/bin\" >> $HOME/.profile", goRoot, goPath)
	if _, err := runCommand("", cmd); err != nil {
		return fmt.Errorf("install: command failed with err: %+v", err)
	}
	fmt.Println(" done")

	fmt.Print(" >>>>>  loading go environment... ")
	_, _ = runCommand("", "source $HOME/.profile")
	fmt.Println(" done")

	fmt.Println("\n Go installation was successful. Run 'go version' to check version installed or run 'go env' to check go-specific environment")
	return nil
}

func installNode(version string) error {
	// configurations for installing node
	var (
		nodeVersion = "10.x"
	)

	if version != "" {
		nodeVersion = fmt.Sprintf("%s.x", version)
	}

	// From Node's official NodeSource Github Page:
	// Replace with the branch of Node.js or io.js you want to install: node_6.x, node_8.x, node_10.x setc...
	// VERSION=node_10.x

	// DISTRO="$(lsb_release -s -c)"
	// echo "deb https://deb.nodesource.com/$VERSION $DISTRO main" | sudo tee /etc/apt/sources.list.d/nodesource.list
	// echo "deb-src https://deb.nodesource.com/$VERSION $DISTRO main" | sudo tee -a /etc/apt/sources.list.d/nodesource.list
	//
	// Replicating above steps...
	nodeRepo := fmt.Sprintf("node_%s", nodeVersion)
	cmdReadDistro := "lsb_release -s -c"

	var out []byte
	var err error
	fmt.Print(" >>>>>  reading host's distro... ")
	if out, err = runCommand("", cmdReadDistro); err != nil {
		return fmt.Errorf("install: command failed with err: %+v", err)
	}
	distro := strings.Trim(string(out), "\n")

	cmdAddNodeSourceRepo := fmt.Sprintf("echo \"deb https://deb.nodesource.com/%s %s main\" | sudo tee /etc/apt/sources.list.d/nodesource.list && echo \"deb-src https://deb.nodesource.com/%s %s main\" | sudo tee -a /etc/apt/sources.list.d/nodesource.list", nodeRepo, distro, nodeRepo, distro)
	fmt.Println(cmdAddNodeSourceRepo)

	fmt.Print(" >>>>>  adding node's source to apt... ")
	if _, err = runCommand("", cmdAddNodeSourceRepo); err != nil {
		return fmt.Errorf("install: command failed with err: %+v", err)
	}

	fmt.Print(" >>>>>  installing node... ")
	cmdInstallNode := "sudo apt-get update && sudo apt-get install nodejs"
	if _, err = runCommand("", cmdInstallNode); err != nil {
		return fmt.Errorf("install: command failed with err: %+v", err)
	}

	return nil
}

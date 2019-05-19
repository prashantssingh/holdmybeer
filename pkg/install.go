package pkg

import "fmt"

// RunInstaller installs a language or a framework based on the flag-value provided to hmb while executing
// the binary on the terminal.
func RunInstaller() error {
	return nil
}

// PrintInstallerUsuage prints all the accepted flags and their respective shorthand letters
func PrintInstallerUsuage() {
	fmt.Println("Usage: hmb install [OPTIONS]")
	fmt.Println()
	fmt.Println("Install specified languages, frameworks and tools")
	fmt.Println()
	fmt.Println("Options: ")
	fmt.Println("  -l, --lang 		Specify a lang to install (supported languages: go, nodejs, java)")
}
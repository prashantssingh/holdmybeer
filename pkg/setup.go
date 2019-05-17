package pkg

import "fmt"

// RunSetup install set of predefined tools based on flag provided to hmb while executing the binary
// on the terminal.
func RunSetup() error {
	return nil
}

// PrintSetupUsuage prints all the accepted flags and their respective shorthand letters
func PrintSetupUsuage() {
	fmt.Println("Usage: hmb setup [OPTIONS]")
	fmt.Println()
	fmt.Println("Setup helps quickly bootstrap an environment")
	fmt.Println()
	fmt.Println("Option: ")
	fmt.Println("  -b, --bare-minimum		Installs tools like vim, curl and htop.")
	fmt.Println("\t\t\t\tUse -i or --ignore flag after this flag to ignore installation of tools you don't need. Refer usuage of -i flag below")
	fmt.Println("  -i, --ignore		Pass a list of comma-separated names of tools enlisted in -b flag, you don't want hmb to install")
}

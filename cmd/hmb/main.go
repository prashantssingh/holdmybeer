package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	cmdInstall = flag.NewFlagSet("install", flag.ExitOnError)
	cmdSetup   = flag.NewFlagSet("setup", flag.ExitOnError)

	flagLang = cmdSetup.StringP("lang", "l", "", "Specify a lang to install. For example: go, nodejs, java")

	flagBareMinimum = cmdSetup.StringP("bare-minimum", "b", "", "Setup bare minimum development environment. This will install vim, curl and htop on host machine")
)

func main() {
	if len(os.Args) == 1 {
		printUsuage()
		return
	}

	if err := run(os.Args[1:]); err != nil {
		log.Fatalf("hmb: command failed to run, err: %+v.\n", err)
		os.Exit(1)
	}
}

func run(cmdArgs []string) error {
	switch cmdArgs[1] {
	case "setup":
		cmdSetup.Parse(cmdArgs[2:])
	case "install":
		cmdInstall.Parse(cmdArgs[2:])
	default:
		log.Fatalf("%q is not a valid command.\n", cmdArgs[1])
		os.Exit(1)
	}

	return nil
}

func printUsuage() {
	fmt.Println("Usage: hmb <command> [<args>]")
	fmt.Println()
	fmt.Println("An installer to aid you with your setup.")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  setup		Setup helps quickly bootstrap an environment")
	fmt.Println("  install		Intall specified languages, frameworks and tools")
}

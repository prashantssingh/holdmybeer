package main

import (
	"fmt"
	"log"
	"os"

	"github.com/prashantssingh/holdmybeer/pkg"

	flag "github.com/spf13/pflag"
)

var (
	cmdSetup   = flag.NewFlagSet("setup", flag.ExitOnError)
	cmdInstall = flag.NewFlagSet("install", flag.ExitOnError)

	flagBareMinimum = cmdSetup.StringP("bare-minimum", "b", "", "Setup bare minimum development environment. This will install vim, curl and htop on host machine")

	flagLang = cmdInstall.StringP("lang", "l", "", "Specify a lang to install. For example: go, nodejs, java")
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
	switch cmdArgs[0] {
	case "setup":
		cmdSetup.Parse(cmdArgs[0:])
	case "install":
		cmdInstall.Parse(cmdArgs[0:])
	default:
		log.Fatalf("%q is not a valid command.\n", cmdArgs[1])
		os.Exit(1)
	}

	if cmdSetup.Parsed() {
		if len(cmdArgs) == 1 || cmdArgs[1] == "-h" || cmdArgs[1] == "--help" {
			pkg.PrintSetupUsuage()
			return nil
		}
		return nil
	}

	if cmdInstall.Parsed() {
		if len(cmdArgs) == 1 || cmdArgs[1] == "-h" || cmdArgs[1] == "--help" {
			pkg.PrintInstallerUsuage()
			return nil
		}
		return nil
	}

	return nil
}

func printUsuage() {
	fmt.Println("Usage: hmb <command> [<args>]")
	fmt.Println()
	fmt.Println("An installer to aid you with your setup")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  setup		Setup helps quickly bootstrap an environment")
	fmt.Println("  install		Install specified languages, frameworks and tools")
}

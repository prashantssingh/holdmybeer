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

	flagBareMinimum = cmdSetup.BoolP("bare-minimum", "b", false, "Setup bare minimum development environment. This will install vim, curl and htop on host machine")

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
			printSetupUsuage()
			return nil
		}

		switch {
		case *flagBareMinimum:
			if err := pkg.SetupBareMinimum(); err != nil {
				log.Fatalf("err: %+v\n", err)
				os.Exit(1)
			}

		}
		return nil
	}

	if cmdInstall.Parsed() {
		if len(cmdArgs) == 1 || cmdArgs[1] == "-h" || cmdArgs[1] == "--help" {
			printInstallerUsuage()
			return nil
		}
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

func printSetupUsuage() {
	fmt.Println("Usage: hmb setup [OPTIONS]")
	fmt.Println()
	fmt.Println("Setup helps quickly bootstrap an environment")
	fmt.Println()
	fmt.Println("Option: ")
	fmt.Println("  -b, --bare-minimum	Installs tools like vim, curl and htop.")
	fmt.Println("\t\t\tUse -i or --ignore flag after this flag to ignore installation of tools you don't need. Refer usuage of -i flag below")
	fmt.Println("  -i, --ignore		Pass a list of comma-separated names of tools enlisted in -b flag to ignore")
}

func printInstallerUsuage() {
	fmt.Println("Usage: hmb install [OPTIONS]")
	fmt.Println()
	fmt.Println("Install specified languages, frameworks and tools")
	fmt.Println()
	fmt.Println("Options: ")
	fmt.Println("  -l, --lang 		Specify a lang to install (supported languages: go, nodejs, java)")
}

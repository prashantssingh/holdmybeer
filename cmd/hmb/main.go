package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cmdInstall = flag.NewFlagSet("install", flag.ExitOnError)

	basicSetup = cmdInstall.String("basic-setup", "", "Setup basic development environment. This will install vim, curl and htop on host machine")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "install":
		cmdInstall.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

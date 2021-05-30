package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func BuildHelpCommand() *Command {

	flagSet := flag.NewFlagSet("help", flag.ExitOnError)
	return &Command{
		Name:         "help",
		FlagSet:      flagSet,
		UsageCommand: "reflorest help <COMMAND>",
		Usage: []string{
			"Print usage information.  If a command is passed in, print usage information just for that command.",
		},
		Command: func(args []string, additionalArgs []string) {
			if len(args) == 0 {
				usage()
			} else {
				command, found := commandMatching(args[0])
				if !found {
					fmt.Printf("Unknown command: %s\n", args[0])
					return
				}
				usageForCommand(command, true)
			}

		},
	}
}

func commandMatching(name string) (*Command, bool) {
	for _, command := range Commands {
		if command.Matches(name) {
			return command, true
		}
	}
	return nil, false
}

func usage() {
	fmt.Fprintf(os.Stderr, "Reflorest Version %s\n\n", VERSION)
	usageForCommand(DefaultCommand, false)
	for _, command := range Commands {
		fmt.Fprintf(os.Stderr, "\n")
		usageForCommand(command, false)
	}
}

func usageForCommand(command *Command, longForm bool) {
	fmt.Fprintf(os.Stderr, "%s\n%s\n", command.UsageCommand, strings.Repeat("-", len(command.UsageCommand)))
	fmt.Fprintf(os.Stderr, "%s\n", strings.Join(command.Usage, "\n"))
	if command.SuppressFlagDocumentation && !longForm {
		fmt.Fprintf(os.Stderr, "%s\n", strings.Join(command.FlagDocSubstitute, "\n  "))
	} else {
		command.FlagSet.PrintDefaults()
	}
}

package main

import (
	"flag"
	"os"
)

const VERSION = "1.0.0"

const greenColor = "\x1b[32m => "
const redColor = "\x1b[91m ** "
const defaultStyle = "\x1b[0m"
const lightGrayColor = "\x1b[37m"

type Command struct {
	Name                      string
	AltName                   string
	FlagSet                   *flag.FlagSet
	Usage                     []string
	UsageCommand              string
	Command                   func(args []string, additionalArgs []string)
	SuppressFlagDocumentation bool
	FlagDocSubstitute         []string
}

func (c *Command) Matches(name string) bool {
	return c.Name == name || (c.AltName != "" && c.AltName == name)
}

func (c *Command) Run(args []string, additionalArgs []string) {
	c.FlagSet.Parse(args)
	c.Command(c.FlagSet.Args(), additionalArgs)
}

var DefaultCommand *Command
var Commands []*Command

func init() {
	DefaultCommand = BuildHelpCommand()
	Commands = append(Commands, BuildBootstrapCommand())
	Commands = append(Commands, BuildDeployCommand())
	Commands = append(Commands, BuildRunCommand())
	Commands = append(Commands, BuildHelpCommand())
}

func main() {
	args := []string{}
	additionalArgs := []string{}

	foundDelimiter := false

	for _, arg := range os.Args[1:] {
		if !foundDelimiter {
			if arg == "--" {
				foundDelimiter = true
				continue
			}
		}

		if foundDelimiter {
			additionalArgs = append(additionalArgs, arg)
		} else {
			args = append(args, arg)
		}
	}

	if len(args) > 0 {
		commandToRun, found := commandMatching(args[0])
		if found {
			commandToRun.Run(args[1:], additionalArgs)
			return
		}
	}

	DefaultCommand.Run(args, additionalArgs)
}

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func BuildRunCommand() *Command {

	flagSet := flag.NewFlagSet("run", flag.ExitOnError)
	return &Command{
		Name:         "run",
		FlagSet:      flagSet,
		UsageCommand: "reflorest run",
		Usage: []string{
			"Run the Application. generally for developers for quick application startup",
		},
		Command: func(args []string, additionalArgs []string) {
			err := run()
			if err != nil {
				fmt.Printf(redColor+"\nError Occured : %s\n"+defaultStyle, err.Error())
				os.Exit(1)
				//panic(err.Error())
			}
		},
	}
}

func run() error {

	fmt.Println(greenColor + "Run Command Started ..." + defaultStyle)

	err := GoRun()
	if err != nil {
		return err
	}
	fmt.Println(greenColor + "Run  Command Successfully Finished" + defaultStyle)
	return nil
}

func GoRun() error {
	return exec.Command("go", "run", "main.go").Run()
}

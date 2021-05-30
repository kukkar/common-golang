package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func BuildDeployCommand() *Command {

	flagSet := flag.NewFlagSet("deploy", flag.ExitOnError)
	return &Command{
		Name:         "deploy",
		FlagSet:      flagSet,
		UsageCommand: "reflorest deploy",
		Usage: []string{
			"Deploy my Application",
		},
		Command: func(args []string, additionalArgs []string) {
			err := deploy()
			if err != nil {
				fmt.Printf(redColor+"\nError Occured : %s\n"+defaultStyle, err.Error())
				os.Exit(1)
			}
		},
	}
}

func deploy() error {

	fmt.Println(greenColor + "Deploy Command Started ..." + defaultStyle)
	err := SetConfigFiles()
	if err != nil {
		return err
	}
	err = GoInstall()
	if err != nil {
		return err
	}
	fmt.Println(greenColor + "Deploy Command Successfully Finished" + defaultStyle)
	return nil
}

func GoInstall() error {
	cmd := exec.Command("go", "install", "./...")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}

func SetConfigFiles() error {
	//move config files to $GOBIN
	var goBinPath string
	var DS = string(os.PathSeparator)
	env := os.Getenv("GOBIN")
	if env == "" {
		goPath := os.Getenv("GOPATH")
		if goPath == "" {
			return fmt.Errorf("$GOPATH and $GOBIN are not set")
		}
		goBinPath = goPath + DS + "bin"
		os.Setenv("GOBIN", goBinPath)
	} else {
		goBinPath = env
	}
	//Remove old config files
	path2conf := goBinPath + DS + "conf"

	//check if folder already exists, if so delete it
	_, err := os.Stat(path2conf)
	if err == nil { // means folder already exists.
		err := os.RemoveAll(path2conf)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}
	fmt.Println(path2conf)
	err = os.MkdirAll(path2conf, 0755)
	if err != nil {
		return err
	}
	//copy config files.
	err = copyFile("conf/conf.json", path2conf+DS+"conf.json")
	if err != nil {
		return err
	}
	err = copyFile("conf/logger.json", path2conf+DS+"logger.json")
	if err != nil {
		return err
	}
	err = copyFile("conf/standard.flf", path2conf+DS+"standard.flf")
	if err != nil {
		return err
	}
	return nil

}

func copyFile(from string, to string) error {
	srcFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = destFile.Sync() //push data to disk
	if err != nil {
		return err
	}
	return nil
}

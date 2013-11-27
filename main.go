package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	log.SetOutput(ioutil.Discard)

	// Get command line arguments.
	args := os.Args[1:]

	cli := &cli.CLI{
		Args:		args,
		Commands:	Commands,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}

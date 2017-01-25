package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/dictav/go-grpcrgen"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

const cmdName = "grpcrgen"

var version = ""

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		dir string

		ver bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(cmdName, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&dir, "framework", "", "Use specific framework [httprouter]")
	flags.BoolVar(&ver, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if ver {
		fmt.Fprintf(cli.errStream, "%s version %s\n", cmdName, version)
		return ExitCodeOK
	}

	if err := grpcrgen.Generate("myservice", "grpcrgen"); err != nil {
		log.Println(err)
		return ExitCodeError
	}

	return ExitCodeOK
}

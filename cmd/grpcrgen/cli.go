package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dictav/go-grpcrgen"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

const (
	cmdName = "grpcrgen"
	version = "v0.1.0"
)

type showVersion struct{}

// implement Error
func (showVersion) Error() string { return "showVersion" }

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
	inputs               []string
	output               string
	framework            string
	ver                  bool
}

func (cli *CLI) parse(args []string) error {
	// Define option flag parse
	flags := flag.NewFlagSet(cmdName, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&cli.output, "o", "router", "Output directory path")
	// flags.StringVar(&cli.framework, "f", "", "Use specific framework. Default is net/http [httprouter]")
	flags.BoolVar(&cli.ver, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	// Show version
	if cli.ver {
		fmt.Fprintln(cli.errStream, version)
		return showVersion{}
	}

	cli.inputs = flags.Args()
	if len(cli.inputs) == 0 {
		flags.Usage()
		return fmt.Errorf("At least one input is required")
	}

	return nil
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {

	if err := cli.parse(args); err != nil {
		if _, ok := err.(showVersion); ok {
			return ExitCodeOK
		}
		return ExitCodeError
	}

	for _, input := range cli.inputs {
		if _, err := os.Stat(input); err != nil {
			fmt.Fprintln(cli.errStream, err.Error())
			return ExitCodeError
		}

		if err := grpcrgen.Generate(input, cli.output); err != nil {
			fmt.Fprintln(cli.errStream, err.Error())
			return ExitCodeError
		}
	}

	return ExitCodeOK
}

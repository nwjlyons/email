package main

import (
	"flag"
	"os"
)

var (
	to      = flag.String("t", "", "Comma separated list of email addresses")
	subject = flag.String("s", "", "Subject")
	body    = flag.String("b", "", "Body")
)

func parseFlags(args []string, file *os.File) {

	// This trick allows command line flags to be be set in unit tests.
	// See https://github.com/VonC/gogitolite/commit/f656a9858cb7e948213db9564a9120097252b429
	a := os.Args[1:]
	if args != nil {
		a = args
	}

	// This enables stdin to be mocked for testing.
	if file != nil {
		stdin = file
	} else {
		stdin = os.Stdin
	}
	flag.CommandLine.Parse(a)
}

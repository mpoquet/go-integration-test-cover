package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/mpoquet/go-integration-test-cover/lib"
	"os"
)

func main() {
	os.Exit(mainReturnWithCode())
}

func mainReturnWithCode() int {
	usage := `Toy program.

Usage:
  toy <TARGET> [options]

Options:
  -h --help     Show this screen.
  --version     Show version.`

	ret := -1
	parser := &docopt.Parser{
		HelpHandler: func(err error, usage string) {
			fmt.Println(usage)
			if err != nil {
				ret = 1
			} else {
				ret = 0
			}
		},
		OptionsFirst: false,
	}

	arguments, _ := parser.ParseArgs(usage, os.Args[1:], "0.1.0")
	if ret != -1 {
		return ret
	}

	lib.Hello(arguments["<TARGET>"].(string))
	return 0
}

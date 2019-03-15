package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRunMain(t *testing.T) {
	var (
		args []string
	)
	for _, arg := range os.Args {
		switch {
		case strings.HasPrefix(arg, "-test"):
		case strings.HasPrefix(arg, "__bypass"):
			args = append(args, strings.TrimPrefix(arg, "__bypass"))
		default:
			args = append(args, arg)
		}
	}
	os.Args = args

	returnCode := mainReturnWithCode()
	fmt.Println("toy return code:", returnCode)
}

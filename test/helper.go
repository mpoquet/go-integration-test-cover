package test

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"
)

func handleCoverage(t *testing.T, expRetCode int) (coverFilename string, expectedReturnCode int) {
	_, exists := os.LookupEnv("DO_COVERAGE")
	if exists {
		coverFilename = t.Name() + ".covout"
		expectedReturnCode = 0
		return
	} else {
		coverFilename = ""
		expectedReturnCode = expRetCode
	}

	return coverFilename, expectedReturnCode
}

func waitCompletionTimeout(completion chan int, timeoutMS int) (exitCode int, err error) {
	select {
	case exitCode := <-completion:
		return exitCode, nil
	case <-time.After(time.Duration(timeoutMS) * time.Millisecond):
		return -1, fmt.Errorf("Timeout reached")
	}
}

func waitOutputTimeout(re *regexp.Regexp, output chan string,
	timeoutMS int, leaveOnNonMatch bool) (matchingLine string, err error) {
	timeoutReached := make(chan int)
	stopTimeout := make(chan int)
	defer close(timeoutReached)
	defer close(stopTimeout)
	go func() {
		select {
		case <-stopTimeout:
		case <-time.After(time.Duration(timeoutMS) * time.Millisecond):
			timeoutReached <- 0
		}
	}()

	for {
		select {
		case line := <-output:
			if re.MatchString(line) {
				stopTimeout <- 0
				return line, nil
			} else {
				if leaveOnNonMatch {
					stopTimeout <- 0
					return line, fmt.Errorf("Non-matching line read: %v", line)
				}
			}
		case <-timeoutReached:
			return "", fmt.Errorf("Timeout reached")
		}
	}
}

func killallToySIGKILL() error {
	cmd := exec.Command("killall")
	cmd.Args = []string{"killall", "-KILL", "--quiet", "toy", "toy.cover"}
	return cmd.Run()
}

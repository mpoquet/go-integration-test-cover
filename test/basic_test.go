package test

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestNoArgs(t *testing.T) {
	args := []string{}
	coverFile, expectedExitCode := handleCoverage(t, 1)

	proc, err := runToyCover(coverFile, args)
	assert.NoError(t, err, "Cannot start toy")
	defer killallToySIGKILL()

	exitCode, err := waitCompletionTimeout(proc.completion, 1000)
	assert.NoError(t, err, "Toy did not complete in time")
	assert.Equal(t, expectedExitCode, exitCode)
}

func TestHelp(t *testing.T) {
	args := []string{"--help"}
	coverFile, expectedExitCode := handleCoverage(t, 0)

	proc, err := runToyCover(coverFile, args)
	assert.NoError(t, err, "Cannot start toy")
	defer killallToySIGKILL()

	exitCode, err := waitCompletionTimeout(proc.completion, 1000)
	assert.NoError(t, err, "Toy did not complete in time")
	assert.Equal(t, expectedExitCode, exitCode)
}

func TestWorld(t *testing.T) {
	args := []string{"world"}
	coverFile, expectedExitCode := handleCoverage(t, 0)

	proc, err := runToyCover(coverFile, args)
	assert.NoError(t, err, "Cannot start toy")
	defer killallToySIGKILL()

	_, err = waitOutputTimeout(regexp.MustCompile(`Hello world!`), proc.outputControl, 1000, true)
	assert.NoError(t, err, "Cannot read 'Hello world!' on toy output")

	exitCode, err := waitCompletionTimeout(proc.completion, 1000)
	assert.NoError(t, err, "Toy did not complete in time")
	assert.Equal(t, expectedExitCode, exitCode)
}

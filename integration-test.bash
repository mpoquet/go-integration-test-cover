#!/usr/bin/env bash
set -eu
base_dir=$(realpath $(dirname $(realpath $0)))
export PATH="${PATH}:${base_dir}"

# Run the tests normally
GOCACHE=off go test -v ${base_dir}/test

# Run the tests with coverage
DO_COVERAGE=1 GOCACHE=off go test -v ${base_dir}/test

# Merge the various coverage files together
gocovmerge ${base_dir}/test/*.covout > ${base_dir}/test/merged.covout

# Get a coverage report. Here, just a text summary.
go tool cover -func ${base_dir}/test/merged.covout

#!/usr/bin/env bash

# This script is in bash to keep it straightforward,
# not because I recommend to compile things directly in bash.
# For a recommendation: https://ninja-build.org/

# Build the normal variant program
go build -o ./toy ./cmd/toy

# Build the cover variant of the program
go test -c -o ./toy.cover -covermode=count -coverpkg=./lib,./cmd/toy ./cmd/toy

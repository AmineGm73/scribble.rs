#!/bin/bash

# Remove the executable if it exists
rm -f ./scribblers

# Build the Go program
go build ./cmd/scribblers

# Run the program
./scribblers
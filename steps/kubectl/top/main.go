package main

import (
	"fmt"
	"os"

	"github.com/stackpulse/steps/kubectl/base"
	"github.com/stackpulse/steps/kubectl/base/top"
)

func run() (int, error) {
	topGet, err := top.NewTop(nil)
	if err != nil {
		return 1, err
	}

	return base.Run(topGet)
}

func main() {
	exitCode, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing top step: %v", err)
	}

	os.Exit(exitCode)
}

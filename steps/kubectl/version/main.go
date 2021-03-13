package main

import (
	"fmt"
	"github.com/stackpulse/public-steps/kubectl/base"
	"github.com/stackpulse/public-steps/kubectl/base/version/get"
	"os"
)

func run() (int, error) {
	versionGet, err := get.NewGetVersion(nil)
	if err != nil {
		return 1, err
	}

	return base.Run(versionGet)
}

func main() {
	exitCode, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing version step: %v", err)
	}

	os.Exit(exitCode)
}

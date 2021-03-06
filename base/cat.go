package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/stackpulse/steps-sdk-go/gcs"
)

type catArgs struct {
	URL string
}

var gCatArgs = &catArgs{}

func defCatArgs(flagSet *flag.FlagSet) {
	flagSet.StringVar(&gCatArgs.URL, "url", "", "GCS object url in the format of: gs://<bucket>/<object path>")
}

func runCat() error {
	bucket, obj, err := gcs.Parse(gCatArgs.URL)
	if err != nil {
		return err
	}

	downloaded, err := gcs.Download(bucket, obj)
	if err != nil {
		return err
	}

	if _, err := os.Stdout.Write(downloaded); err != nil {
		return fmt.Errorf("can't write to stdout: %w", err)
	}

	return nil
}

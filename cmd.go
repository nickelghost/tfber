package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
)

var errUnknownCommand = errors.New("unknown command")

//go:embed help.txt
var helpText string

func dispatchCmd(cmd string, args []string) error {
	var err error

	switch cmd {
	case "generate":
		err = generateCmd(args)
	case "output":
		err = outputCmd(args)
	case "import":
		err = importCmd(args)
	default:
		fmt.Fprint(os.Stderr, helpText)

		return errUnknownCommand
	}

	return err
}

func initProvider(f flags) (provider, error) {
	switch f.provider {
	case "aws":
		return &awsProvider{}, nil
	case "gcp":
		return &gcpProvider{bucketLocation: f.gcpBucketLocation, project: f.gcpProject}, nil
	}

	return nil, fmt.Errorf("%w: %s", errUnknownProvider, f.provider)
}

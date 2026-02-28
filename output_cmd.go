package main

import (
	"flag"
	"fmt"
)

func outputCmd(args []string) error {
	fs := flag.NewFlagSet("output", flag.ContinueOnError)

	f := flags{}
	f.registerCommon(fs)
	f.registerOutput(fs)

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("%w: %w", errFlagParseFailed, err)
	}

	if err := f.validateCommon(fs); err != nil {
		return err
	}

	if err := f.validateGCP(fs); err != nil {
		return err
	}

	p, err := initProvider(f)
	if err != nil {
		return err
	}

	tfCode := output(p, f.resourceName, f.resourceNameSuffix, f.resourceNameSuffixLength, f.stateID)

	fmt.Print(tfCode) //nolint:forbidigo

	return nil
}

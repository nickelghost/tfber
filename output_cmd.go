package main

import (
	"flag"
	"fmt"
)

func outputCmd(args []string) error {
	fs := flag.NewFlagSet("output", flag.ExitOnError)

	f := flags{}
	f.registerCommon(fs)
	f.registerOutput(fs)

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("flag set parse: %w", err)
	}

	if err := f.validateCommon(fs); err != nil {
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

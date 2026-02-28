package main

import (
	"flag"
	"fmt"
)

func importCmd(args []string) error {
	fs := flag.NewFlagSet("import", flag.ContinueOnError)

	f := flags{}
	f.registerCommon(fs)

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("%w: %w", errFlagParseFailed, err)
	}

	if err := f.validateCommon(fs); err != nil {
		return err
	}

	if err := f.validateImport(fs); err != nil {
		return err
	}

	p, err := initProvider(f)
	if err != nil {
		return err
	}

	return importExec(p, f.resourceName, f.stateID)
}

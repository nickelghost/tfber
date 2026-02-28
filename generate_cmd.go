package main

import (
	"flag"
	"fmt"
)

func generateCmd(args []string) error {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)

	f := flags{}
	f.registerCommon(fs)
	f.registerOutput(fs)

	fileName := fs.String("file-name", "state.tf", "name the generated Terraform file should have")

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

	return generate(p, f.resourceName, f.resourceNameSuffix, f.resourceNameSuffixLength, f.stateID, *fileName)
}

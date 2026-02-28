package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func generateCmd(args []string) error {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)

	f := flags{}
	f.registerCommon(fs)
	f.registerOutput(fs)

	fileName := fs.String("file-name", "state.tf", "name the generated Terraform file should have")
	force := fs.Bool("force", false, "overwrite the file if it already exists")

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

	err = generate(p, f.resourceName, f.resourceNameSuffix, f.resourceNameSuffixLength, f.stateID, *fileName, *force)
	if errors.Is(err, errFileExists) {
		fmt.Fprintf(os.Stderr, "file %s already exists, use -force to overwrite\n", *fileName)
	}

	return err
}

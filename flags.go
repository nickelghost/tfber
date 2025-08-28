package main

import (
	"errors"
	"flag"
	"fmt"
)

var errFlagValidationFailed = errors.New("flag validation failed")

type flags struct {
	// common
	provider     string
	resourceName string
	stateID      string
	// output
	resourceNameSuffix       bool
	resourceNameSuffixLength int
	// output gcp
	gcpBucketLocation string
	gcpProject        string
}

func (f *flags) registerCommon(fs *flag.FlagSet) {
	fs.StringVar(&f.provider, "provider", "", "name of the state provider (required)")
	fs.StringVar(&f.resourceName, "resource-name", "", "name of the resources inside provider (required)")
	fs.StringVar(&f.stateID, "state-id", "", "ID of the resources inside of Terraform path (required on import)")
}

func (f *flags) registerOutput(fs *flag.FlagSet) {
	fs.BoolVar(
		&f.resourceNameSuffix,
		"resource-name-suffix",
		true,
		"name of the resource should have a random suffix",
	)
	fs.IntVar(
		&f.resourceNameSuffixLength,
		"resource-name-suffix-length",
		0,
		"how many characters long the suffix should be (if resource-name-suffix enabled)",
	)
	fs.StringVar(&f.gcpBucketLocation, "gcp-bucket-location", "", "where the GCP bucket should live")
	fs.StringVar(&f.gcpProject, "gcp-project", "", "what project the resources should be created in")
}

func (f *flags) validateCommon(fs *flag.FlagSet) error {
	if f.provider == "" {
		fs.Usage()

		return fmt.Errorf("%w: missing provider", errFlagValidationFailed)
	}

	if f.resourceName == "" {
		fs.Usage()

		return fmt.Errorf("%w: missing resource name", errFlagValidationFailed)
	}

	return nil
}

func (f *flags) validateImport(fs *flag.FlagSet) error {
	if f.stateID == "" {
		fs.Usage()

		return fmt.Errorf("%w: missing state ID", errFlagValidationFailed)
	}

	return nil
}

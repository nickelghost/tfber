package main

import (
	"fmt"
	"os"
)

func generate(
	p provider,
	resourceName string,
	resourceNameSuffix bool,
	resourceNameSuffixLength int,
	stateID, fileName string,
) error {
	code := output(p, resourceName, resourceNameSuffix, resourceNameSuffixLength, stateID)

	f, err := os.Create(fileName) //nolint:gosec
	if err != nil {
		return fmt.Errorf("create file %s: %w", fileName, err)
	}

	defer func() { _ = f.Close() }()

	_, err = f.WriteString(code)
	if err != nil {
		return fmt.Errorf("write to file %s: %w", fileName, err)
	}

	return nil
}

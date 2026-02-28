package main

import (
	"errors"
	"fmt"
	"os"
)

var errFileExists = errors.New("file already exists")

func generate(
	p provider,
	resourceName string,
	resourceNameSuffix bool,
	resourceNameSuffixLength int,
	stateID, fileName string,
	force bool,
) error {
	if !force {
		if _, err := os.Stat(fileName); err == nil {
			return fmt.Errorf("%s: %w", fileName, errFileExists)
		} else if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("stat file %s: %w", fileName, err)
		}
	}

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

package main

import "fmt"

func importExec(p provider, resourceName, stateID string) error {
	if err := p.Import(resourceName, stateID); err != nil {
		return fmt.Errorf("provider import functionality: %w", err)
	}

	return nil
}

package main

import (
	"context"
	"fmt"
)

func importExec(p provider, resourceName, stateID string) error {
	ctx := context.Background()

	if err := p.Import(ctx, resourceName, stateID); err != nil {
		return fmt.Errorf("provider import functionality: %w", err)
	}

	return nil
}

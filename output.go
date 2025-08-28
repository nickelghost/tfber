package main

import "strings"

func output(
	p provider, resourceName string, resourceNameSuffix bool, resourceNameSuffixLength int, stateID string,
) string {
	if stateID == "" {
		stateID = strings.ReplaceAll(resourceName, "-", "_")
	}

	if resourceNameSuffix {
		resourceName += p.CreateResourceNameSuffix(resourceNameSuffixLength)
	}

	return p.Output(resourceName, stateID)
}

package main

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
)

const defaultSuffixLength = 6

var errUnknownProvider = errors.New("unknown provider")

type provider interface {
	Output(resourceName, stateID string) string
	CreateResourceNameSuffix(length int) string
	Import(ctx context.Context, resourceName, stateID string) error
}

func createGenericSuffix(length int) string {
	if length == 0 {
		length = defaultSuffixLength
	}

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, length)

	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		b[i] = chars[n.Int64()]
	}

	return "-" + string(b)
}

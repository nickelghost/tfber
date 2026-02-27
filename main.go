package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	cmd := strings.ToLower(flag.Arg(0))
	args := flag.Args()

	if len(args) > 0 {
		args = args[1:]
	}

	if err := dispatchCmd(cmd, args); err != nil {
		if !errors.Is(err, errFlagValidationFailed) && !errors.Is(err, errUnknownCommand) {
			fmt.Fprintln(os.Stderr, err)
		}

		os.Exit(1)
	}
}

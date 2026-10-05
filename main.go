// Package main is the entrypoint to a REPL for the Monkey Language.
package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/NeilMcB/go-compiler/repl"
)

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(in io.Reader, out io.Writer) error {
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("looking up current user: %w", err)
	}

	_, err = fmt.Fprintf(out, "Hello, %s! This is the Monkey programming language!\n", u.Username)
	if err != nil {
		return fmt.Errorf("writing greeting: %w", err)
	}

	_, err = fmt.Fprint(out, "Feel free to type in some commands\n")
	if err != nil {
		return fmt.Errorf("writing greeting: %w", err)
	}

	return repl.Start(in, out)
}

// Package repl implements a read-eval-print-loop for our language.
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/NeilMcB/go-compiler/lexer"
	"github.com/NeilMcB/go-compiler/token"
)

const prompt = ">> "

// Start begins a new REPL.
func Start(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprint(out, prompt); err != nil {
			return fmt.Errorf("writing prompt: %w", err)
		}

		if !scanner.Scan() {
			break
		}

		if err := writeTokens(scanner.Text(), out); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	return nil
}

func writeTokens(line string, out io.Writer) error {
	l := lexer.New(line)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		if _, err := fmt.Fprintf(out, "%+v\n", tok); err != nil {
			return fmt.Errorf("writing token: %w", err)
		}
	}

	return nil
}

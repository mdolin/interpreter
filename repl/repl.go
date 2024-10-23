package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// Initializes tok with l.NextToken().
		// Continues looping while tok.Type is not token.EOF.
		// Updates tok with l.NextToken() at the end of each iteration.
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { // For loop with initialization, condition, and post statements all in one line
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

package repl

/*
REPL stands for Read-Eval-Print Loop. It is an interactive programming environment that allows users to:

Read: Accept input from the user (code or commands).
Evaluate: Execute or process the input.
Print: Display the result of the evaluation to the user.
Loop: Repeat the process, allowing the user to continue interacting.

e.g., Python: The Python interpreter (python or python3 command).
*/

import (
	"bufio"
	"fmt"
	"io"
	"ravs_lang/lexer"
	"ravs_lang/token"
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

		if line == "exit()" {
			break
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

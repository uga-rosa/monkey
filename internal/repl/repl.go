package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/uga-rosa/monkey/internal/lexer"
	"github.com/uga-rosa/monkey/internal/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		fmt.Fprint(out, PROMPT)

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

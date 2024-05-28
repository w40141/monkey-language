package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/w40141/monkey-language/golang/lexer"
	"github.com/w40141/monkey-language/golang/token"
)

// PROMPT is the repl prompt
const PROMPT = ">> "

// Start starts the REPL
func Start(in io.Reader, _ io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

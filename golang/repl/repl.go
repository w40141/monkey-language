// Package repl implements the Read-Eval-Print-Loop for the Monkey programming language.
package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/w40141/monkey-language/golang/evaluator"
	"github.com/w40141/monkey-language/golang/lexer"
	"github.com/w40141/monkey-language/golang/parser"
)

// prompt is the repl prompt
const (
	prompt     = ">> "
	monkeyFace = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`
)

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		prg := p.ParseProgram()
		if len(p.Errors()) != 0 {
			if e := printParserErrors(out, p.Errors()); e != nil {
				log.Fatal(e)
			}
			continue
		}

		evaluated := evaluator.Eval(prg)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) error {
	io.WriteString(out, monkeyFace)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		if _, e := io.WriteString(out, "\t"+msg+"\n"); e != nil {
			return e
		}
	}
	return nil
}

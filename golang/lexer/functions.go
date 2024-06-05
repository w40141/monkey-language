// Package lexer implements a lexer for the Monkey programming language.
package lexer

// New returns a new instance of Lexer.
func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

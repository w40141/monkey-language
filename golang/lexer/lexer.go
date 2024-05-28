package lexer

import (
	"github.com/w40141/monkey-language/golang/token"
)

// Lexer represents a lexer.
type Lexer struct {
	input string
	// current position in input (points to current char)
	position int
	// current reading position in input (after current char)
	readPosition int
	// current char under examination
	ch byte
}

// New returns a new instance of Lexer.
func New(input string) Lexer {
	l := Lexer{input: input}
	newLexer := l.readChar()
	return newLexer
}

func (l Lexer) readChar() Lexer {
	newLexer := Lexer{input: l.input}
	if l.readPosition >= len(l.input) {
		newLexer.ch = 0
	} else {
		newLexer.ch = l.input[l.readPosition]
	}
	newLexer.position = l.readPosition
	newLexer.readPosition = l.readPosition + 1
	return newLexer
}

// NextToken returns the next token.
func (l Lexer) NextToken() (Lexer, token.Token) {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '*':
		tok = token.New(token.TIMES, l.ch)
	case '/':
		tok = token.New(token.DIVIDE, l.ch)
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '(':
		tok = token.New(token.LPARAN, l.ch)
	case ')':
		tok = token.New(token.RPARAN, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			var nl Lexer
			nl, tok.Literal = l.readIdentifier()
			return nl, tok
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}
	newLexer := l.readChar()
	return newLexer, tok
}

func (l Lexer) readIdentifier() (Lexer, string) {
	position := l.position
	nl := l
	for isLetter(nl.ch) {
		nl = nl.readChar()
	}
	return nl, l.input[position:nl.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

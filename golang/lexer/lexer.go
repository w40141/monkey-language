// Package lexer implements a lexer for the Monkey programming language.
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

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.New(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '*':
		tok = token.New(token.TIMES, l.ch)
	case '/':
		tok = token.New(token.DIVIDE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.NQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.New(token.BANG, l.ch)
		}
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
	case '<':
		tok = token.New(token.LT, l.ch)
	case '>':
		tok = token.New(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	nl := l
	for isLetter(nl.ch) {
		l.readChar()
	}
	return l.input[position:nl.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	nl := l
	for isDigit(nl.ch) {
		l.readChar()
	}
	return l.input[position:nl.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

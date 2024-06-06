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
	nowChar byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 means EOF
		l.nowChar = 0
	} else {
		l.nowChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.nowChar {
	case '=':
		if l.peekChar() == '=' {
			ch := l.nowChar
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(l.nowChar),
			}
		} else {
			tok = token.New(token.ASSIGN, l.nowChar)
		}
	case '+':
		tok = token.New(token.PLUS, l.nowChar)
	case '-':
		tok = token.New(token.MINUS, l.nowChar)
	case '*':
		tok = token.New(token.TIMES, l.nowChar)
	case '/':
		tok = token.New(token.DIVIDE, l.nowChar)
	case '!':
		if l.peekChar() == '=' {
			ch := l.nowChar
			l.readChar()
			tok = token.Token{
				Type:    token.NQ,
				Literal: string(ch) + string(l.nowChar),
			}
		} else {
			tok = token.New(token.BANG, l.nowChar)
		}
	case ',':
		tok = token.New(token.COMMA, l.nowChar)
	case ';':
		tok = token.New(token.SEMICOLON, l.nowChar)
	case '(':
		tok = token.New(token.LPARAN, l.nowChar)
	case ')':
		tok = token.New(token.RPARAN, l.nowChar)
	case '{':
		tok = token.New(token.LBRACE, l.nowChar)
	case '}':
		tok = token.New(token.RBRACE, l.nowChar)
	case '<':
		tok = token.New(token.LT, l.nowChar)
	case '>':
		tok = token.New(token.GT, l.nowChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.nowChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.nowChar) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.nowChar)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	nl := l
	for isLetter(nl.nowChar) {
		l.readChar()
	}
	return l.input[position:nl.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	nl := l
	for isDigit(nl.nowChar) {
		l.readChar()
	}
	return l.input[position:nl.position]
}

func (l *Lexer) skipWhitespace() {
	for l.nowChar == ' ' || l.nowChar == '\t' || l.nowChar == '\n' || l.nowChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

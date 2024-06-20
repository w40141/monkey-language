// Package token defines constants representing the lexical tokens.
package token

const (
	// ILLEGAL represent Special tokens
	ILLEGAL = "ILLEGAL"
	// EOF represents end of the file.
	EOF = "EOF"

	// IDENT represents identifiers.
	IDENT = "IDENT"
	// INT represents integers.
	INT = "INT"
	// FLOAT represents floating point numbers.
	FLOAT = "FLOAT"
	// STRING represents strings.
	STRING = "STRING"

	// ASSIGN represents assignment operator.
	ASSIGN = "="
	// PLUS represents addition operator.
	PLUS = "+"
	// MINUS represents subtraction operator.
	MINUS = "-"
	// TIMES represents multiplication operator.
	TIMES = "*"
	// DIVIDE represents division operator.
	DIVIDE = "/"
	// BANG represents bang operator.
	BANG = "!"

	// COMMA represents comma.
	COMMA = ","
	// SEMICOLON represents semicolon.
	SEMICOLON = ";"

	// LPARAN represents left parenthesis.
	LPARAN = "("
	// RPARAN represents right parenthesis.
	RPARAN = ")"
	// LBRACE represents left brace.
	LBRACE = "{"
	// RBRACE represents right brace.
	RBRACE = "}"
	// LBRACKET represents left bracket.
	LBRACKET = "["
	// RBRACKET represents right bracket.
	RBRACKET = "]"
	// LT represents less than.
	LT = "<"
	// GT represents greater than.
	GT = ">"

	// FUNCTION represents function keyword.
	FUNCTION = "FUNCTION"
	// LET represents let keyword.
	LET = "LET"
	// TRUE represents true keyword.
	TRUE = "TRUE"
	// FALSE represents false keyword.
	FALSE = "FALSE"
	// IF represents if keyword.
	IF = "IF"
	// ELSE represents else keyword.
	ELSE = "ELSE"
	// RETURN represents return keyword.
	RETURN = "RETURN"

	// EQ represents equal operator.
	EQ = "=="
	// NQ represents not equal operator.
	NQ = "!="
)

// Type is the type of token.
type Type string

// Token is a lexical token.
type Token struct {
	Type    Type
	Literal string
}

// New returns a new instance of Token.
func New(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns the token type of the given identifier.
// TODO: test
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

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

	// FUNCTION represents function keyword.
	FUNCTION = "FUNCTION"
	// LET represents let keyword.
	LET = "LET"
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
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent returns the token type of the given identifier.
// TODO: test
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

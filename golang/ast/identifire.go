// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*Identifier)(nil)

// Identifier represents an identifier in the AST.
type Identifier struct {
	// Token is the token.IDENT token.
	Token token.Token
	Value string
}

// String implements Expression.
func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the token literal of the identifier.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

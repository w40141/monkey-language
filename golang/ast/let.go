// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Statement = (*LetStatement)(nil)

// LetStatement represents a let statement in the AST.
type LetStatement struct {
	// Token is the token.LET token.
	Token token.Token
	Name  *Identifier
	Value Expression
}

// String implements Statement.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token literal of the let statement.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

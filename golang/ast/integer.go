// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import "github.com/w40141/monkey-language/golang/token"

var _ Expression = (*IntegerLiteral)(nil)

// IntegerLiteral represents an integer literal in the AST.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// expressionNode implements Expression.
func (i *IntegerLiteral) expressionNode() {}

// String implements Node.
func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

// TokenLiteral implements Node.
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import "github.com/w40141/monkey-language/golang/token"

var _ Expression = (*Boolean)(nil)

// Boolean represents a boolean value in the AST.
type Boolean struct {
	Token token.Token
	Value bool
}

// String implements Expression.
func (b *Boolean) String() string {
	return b.Token.Literal
}

// TokenLiteral implements Expression.
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// expressionNode implements Expression.
func (b *Boolean) expressionNode() {}

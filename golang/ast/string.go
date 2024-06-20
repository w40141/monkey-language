// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import "github.com/w40141/monkey-language/golang/token"

var _ Expression = (*StringLiteral)(nil)

// StringLiteral represents a string literal in the AST.
type StringLiteral struct {
	Token token.Token
	Value string
}

// String implements Expression.
func (s *StringLiteral) String() string {
	return s.Token.Literal
}

// TokenLiteral implements Expression.
func (s *StringLiteral) TokenLiteral() string {
	return s.Token.Literal
}

// expressionNode implements Expression.
func (s *StringLiteral) expressionNode() {
}

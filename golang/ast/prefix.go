// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*PrefixExpression)(nil)

// PrefixExpression represents a prefix expression in the AST.
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// String implements Expression.
func (p *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")
	return out.String()
}

// TokenLiteral implements Expression.
func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

// expressionNode implements Expression.
func (p *PrefixExpression) expressionNode() {}

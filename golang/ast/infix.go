// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*InfixExpression)(nil)

// InfixExpression represents an infix expression in the AST.
type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

// String implements Expression.
func (p *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Left.String())
	out.WriteString(" " + p.Operator + " ")
	out.WriteString(p.Right.String())
	out.WriteString(")")
	return out.String()
}

// TokenLiteral implements Expression.
func (p *InfixExpression) TokenLiteral() string {
	return p.Token.Literal
}

// expressionNode implements Expression.
func (p *InfixExpression) expressionNode() {}

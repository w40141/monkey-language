// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*IfExpression)(nil)

// IfExpression represents an if expression in the AST.
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	ALternative *BlockStatement
}

// String implements Expression.
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.ALternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.ALternative.String())
	}

	return out.String()
}

// TokenLiteral implements Expression.
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// expressionNode implements Expression.
func (ie *IfExpression) expressionNode() {}


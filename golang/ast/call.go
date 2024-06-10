// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"
	"strings"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*CallExpression)(nil)

// CallExpression represents a call expression in the AST.
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

// String implements Expression.
func (c *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, arg := range c.Arguments {
		args = append(args, arg.String())
	}

	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// TokenLiteral implements Expression.
func (c *CallExpression) TokenLiteral() string {
	return c.Token.Literal
}

// expressionNode implements Expression.
func (c *CallExpression) expressionNode() {
}

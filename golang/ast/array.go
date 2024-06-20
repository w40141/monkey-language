// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"
	"strings"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*ArrayLiteral)(nil)

// ArrayLiteral represents an array literal in the AST.
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

// String implements Expression.
func (a *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range a.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// TokenLiteral implements Expression.
func (a *ArrayLiteral) TokenLiteral() string {
	return a.Token.Literal
}

// expressionNode implements Expression.
func (a *ArrayLiteral) expressionNode() {
}

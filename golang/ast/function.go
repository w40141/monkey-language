// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"
	"strings"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*FunctionLiteral)(nil)

// FunctionLiteral represents a function literal in the AST.
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

// String implements Expression.
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// TokenLiteral implements Expression.
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// expressionNode implements Expression.
func (f *FunctionLiteral) expressionNode() {}

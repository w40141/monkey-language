// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"
	"strings"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*HashLiteral)(nil)

// HashLiteral represents a hash literal in the AST.
type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

// String implements Expression.
func (h *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range h.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// TokenLiteral implements Expression.
func (h *HashLiteral) TokenLiteral() string {
	return h.Token.Literal
}

// expressionNode implements Expression.
func (h *HashLiteral) expressionNode() {
}

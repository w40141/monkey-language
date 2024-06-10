// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*BlockStatement)(nil)

// BlockStatement represents a block statement in the AST.
type BlockStatement struct {
	Token     token.Token
	Statements []Statement
}

// String implements Expression.
func (b *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range b.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// TokenLiteral implements Expression.
func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

// expressionNode implements Expression.
func (b *BlockStatement) expressionNode() {}

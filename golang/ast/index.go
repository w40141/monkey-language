// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Expression = (*IndexExpression)(nil)

// IndexExpression represents an index expression in the AST.
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

// String implements Expression.
func (i *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString("[")
	out.WriteString(i.Index.String())
	out.WriteString("])")

	return out.String()
}

// TokenLiteral implements Expression.
func (i *IndexExpression) TokenLiteral() string {
	return i.Token.Literal
}

// expressionNode implements Expression.
func (i *IndexExpression) expressionNode() {
}

// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"bytes"

	"github.com/w40141/monkey-language/golang/token"
)

var _ Statement = (*ReturnStatement)(nil)

// ReturnStatement represents a return statement in the AST.
type ReturnStatement struct {
	// Token is the token.RETURN token.
	Token       token.Token
	ReturnValue Expression
}

// String implements Statement.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the token literal of the return statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

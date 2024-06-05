// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import "github.com/w40141/monkey-language/golang/token"

var _ Statement = (*ExpressionStatement)(nil)

// ExpressionStatement represents an expression statement in the AST.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// String implements Statement.
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the token literal of the expression statement.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

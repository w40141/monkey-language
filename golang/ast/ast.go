// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"github.com/w40141/monkey-language/golang/token"
)

// Node is the interface that all nodes in the AST implement.
type Node interface {
	TokenLiteral() string
}

// Program represents a Monkey program.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the token literal of the first statement in the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Statement is the interface that all statement nodes in the AST implement.
type Statement interface {
	Node
	statementNode()
}

// LetStatement represents a let statement in the AST.
type LetStatement struct {
	// Token is the token.LET token.
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token literal of the let statement.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier represents an identifier in the AST.
type Identifier struct {
	// Token is the token.IDENT token.
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the token literal of the identifier.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Expression is the interface that all expression nodes in the AST implement.
type Expression interface {
	Node
	expressionNode()
}

// ReturnStatement represents a return statement in the AST.
type ReturnStatement struct {
	// Token is the token.RETURN token.
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the token literal of the return statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

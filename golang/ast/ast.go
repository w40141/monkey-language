// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

// Node is the interface that all nodes in the AST implement.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is the interface that all statement nodes in the AST implement.
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface that all expression nodes in the AST implement.
type Expression interface {
	Node
	expressionNode()
}

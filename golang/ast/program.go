// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import "bytes"

var _ Node = (*Program)(nil)

// Program represents a Monkey program.
type Program struct {
	Statements []Statement
}

// String implements Node.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// TokenLiteral returns the token literal of the first statement in the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

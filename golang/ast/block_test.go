// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"testing"

	"github.com/w40141/monkey-language/golang/token"
)

func TestString_Block(t *testing.T) {
	input := &BlockStatement{
		Token: token.Token{Type: token.LBRACE, Literal: "{"},
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &IntegerLiteral{
					Token: token.Token{
						Type:    token.INT,
						Literal: "5",
					},
					Value: 5,
				},
			},
		},
	}
	if input.String() != "let myVar = 5;" {
		t.Errorf("input.String() = %q, want %q", input.String(), "let myVar = 5;")
	}
}

func TestTokenLiteral_Block(t *testing.T) {
	input := &BlockStatement{
		Token: token.Token{Type: token.LBRACE, Literal: "{"},
	}
	if input.TokenLiteral() != "{" {
		t.Errorf("input.TokenLiteral() = %q, want %q", input.TokenLiteral(), "{")
	}
}

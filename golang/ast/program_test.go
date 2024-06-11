// Package ast defines the abstract syntax tree for the Monkey programming language.
package ast

import (
	"testing"

	"github.com/w40141/monkey-language/golang/token"
)

func TestString_Program(t *testing.T) {
	tests := []struct {
		input *Program
		want  string
	}{
		{
			input: &Program{
				Statements: []Statement{
					&LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name: &Identifier{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "myVar",
							},
							Value: "myVar",
						},
						Value: &Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
							Value: "anotherVar",
						},
					},
				},
			},
			want: "let myVar = anotherVar;",
		},
	}

	for _, tt := range tests {
		p := tt.input
		if got := p.String(); got != tt.want {
			t.Errorf("p.String() = %q, want %q", got, tt.want)
		}
	}
}

func TestTokenLiteral_Program(t *testing.T) {
	tests := []struct {
		input *Program
		want  string
	}{
		{
			input: &Program{Statements: []Statement{}},
			want:  "",
		},
		{
			input: &Program{
				Statements: []Statement{
					&LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
					},
				},
			},
			want: "let",
		},
	}

	for _, tt := range tests {
		p := tt.input
		if got := p.TokenLiteral(); got != tt.want {
			t.Errorf("p.TokenLiteral() = %q, want %q", got, tt.want)
		}
	}
}

package token

import "testing"

func TestNew(t *testing.T) {
	type inputs struct {
		tokenType Type
		ch        byte
	}

	tests := []struct {
		input inputs
		want  Token
	}{
		{
			input: inputs{
				tokenType: ASSIGN,
				ch:        '=',
			},
			want: Token{
				Type:    ASSIGN,
				Literal: "=",
			},
		},
		{
			input: inputs{
				tokenType: PLUS,
				ch:        '+',
			},
			want: Token{
				Type:    PLUS,
				Literal: "+",
			},
		},
	}

	for i, tt := range tests {
		tok := New(tt.input.tokenType, tt.input.ch)
		if tok != tt.want {
			t.Fatalf(
				"tests[%d] - token wrong. expected=%+v, got=%+v",
				i,
				tt.want,
				tok,
			)
		}
	}
}

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input string
		want  Type
	}{
		{
			input: "fn",
			want:  FUNCTION,
		},
		{
			input: "let",
			want:  LET,
		},
		// {
		// 	input:  "true",
		// 	want: TRUE,
		// },
		// {
		// 	input:  "false",
		// 	want: FALSE,
		// },
		// {
		// 	input:  "if",
		// 	want: IF,
		// },
		// {
		// 	input:  "else",
		// 	want: ELSE,
		// },
		// {
		// 	input:  "return",
		// 	want: RETURN,
		// },
		{
			input: "foobar",
			want:  IDENT,
		},
	}
	for i, tt := range tests {
		ident := LookupIdent(tt.input)
		if ident != tt.want {
			t.Fatalf(
				"tests[%d] - token wrong. expected=%+v, got=%+v",
				i,
				tt.want,
				ident,
			)
		}
	}
}

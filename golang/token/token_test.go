package token

import "testing"

func TestNew(t *testing.T) {
	type inputs struct {
		tokenType Type
		ch        byte
	}

	tests := []struct {
		name  string
		input inputs
		want  Token
	}{
		{
			name: "ASSIGN token",
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
			name: "SEMICOLON token",
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

	for _, tt := range tests {
		tok := New(tt.input.tokenType, tt.input.ch)
		if tok != tt.want {
			t.Fatalf(
				"name: %s - token wrong. expected=%+v, got=%+v",
				tt.name,
				tt.want,
				tok,
			)
		}
	}
}

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Type
	}{
		{
			name:  "Function",
			input: "fn",
			want:  FUNCTION,
		},
		{
			name:  "Let",
			input: "let",
			want:  LET,
		},
		{
			name:  "True",
			input: "true",
			want:  TRUE,
		},
		{
			name:  "False",
			input: "false",
			want:  FALSE,
		},
		{
			name:  "If",
			input: "if",
			want:  IF,
		},
		{
			name:  "Else",
			input: "else",
			want:  ELSE,
		},
		{
			name:  "Return",
			input: "return",
			want:  RETURN,
		},
		{
			name:  "Ident",
			input: "foobar",
			want:  IDENT,
		},
	}
	for _, tt := range tests {
		ident := LookupIdent(tt.input)
		if ident != tt.want {
			t.Fatalf(
				"name: %s - token wrong. expected=%+v, got=%+v",
				tt.name,
				tt.want,
				ident,
			)
		}
	}
}

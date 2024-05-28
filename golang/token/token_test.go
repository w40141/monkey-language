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

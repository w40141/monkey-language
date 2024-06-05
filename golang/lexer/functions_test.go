package lexer

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input string
		wants *Lexer
	}{
		{
			input: `=`,
			wants: &Lexer{
				input:        "=",
				position:     0,
				readPosition: 1,
				ch:           '=',
			},
		},
		{
			input: ``,
			wants: &Lexer{
				input:        "",
				position:     0,
				readPosition: 1,
				ch:           0,
			},
		},
	}

	for i, tt := range tests {
		l := New(tt.input)
		if !reflect.DeepEqual(l, tt.wants) {
			t.Fatalf(
				"tests[%d] - lexer wrong. expected=%+v, got=%+v",
				i,
				tt.wants,
				l,
			)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		input byte
		want  bool
	}{
		{'a', true},
		{'A', true},
		{'z', true},
		{'Z', true},
		{'_', true},
		{'1', false},
	}

	for i, tt := range tests {
		if got := isLetter(tt.input); got != tt.want {
			t.Fatalf(
				"tests[%d] - isLetter wrong. expected=%t, got=%t",
				i,
				tt.want,
				got,
			)
		}
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input byte
		want  bool
	}{
		{'0', true},
		{'9', true},
		{'a', false},
	}
	for i, tt := range tests {
		if got := isDigit(tt.input); got != tt.want {
			t.Fatalf(
				"tests[%d] - isDigit wrong. expected=%t, got=%t",
				i,
				tt.want,
				got,
			)
		}
	}
}

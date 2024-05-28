package lexer

import (
	"reflect"
	"testing"

	"github.com/w40141/monkey-language/golang/token"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input string
		wants Lexer
	}{
		{
			input: `=`,
			wants: Lexer{
				input:        "=",
				position:     0,
				readPosition: 1,
				ch:           '=',
			},
		},
		{
			input: ``,
			wants: Lexer{
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

func TestReadChar(t *testing.T) {
	tests := []struct {
		input Lexer
		want  Lexer
	}{
		{
			input: Lexer{
				input:        "=",
				position:     0,
				readPosition: 1,
				ch:           '=',
			},
			want: Lexer{
				input:        "=",
				position:     1,
				readPosition: 2,
				ch:           0,
			},
		},
	}

	for i, tt := range tests {
		newLexer := tt.input.readChar()
		if !reflect.DeepEqual(newLexer, tt.want) {
			t.Fatalf(
				"tests[%d] - lexer wrong. expected=%+v, got=%+v",
				i,
				tt.want,
				newLexer,
			)
		}
	}
}

func TestNextToken(t *testing.T) {
	type want struct {
		ttype   token.Type
		literal string
	}

	tests := []struct {
		input string
		wants []want
	}{
		{
			input: `=+(){},;`,
			wants: []want{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPARAN, "("},
				{token.RPARAN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		// {
		// 	input: `
		// 	let five = 5;
		// 	let ten = 10;
		// 	let add = fn(x, y) {
		// 		x + y;
		// 	}
		// 	let result = add(five, ten);
		// 	`,
		// 	wants: []want{
		// 		{token.LET, "let"},
		// 		{token.IDENT, "five"},
		// 		{token.ASSIGN, "="},
		// 		{token.INT, "5"},
		// 		{token.SEMICOLON, ";"},
		// 		{token.LET, "let"},
		// 		{token.IDENT, "ten"},
		// 		{token.ASSIGN, "="},
		// 		{token.INT, "10"},
		// 		{token.SEMICOLON, ";"},
		// 		{token.LET, "let"},
		// 		{token.IDENT, "add"},
		// 		{token.ASSIGN, "="},
		// 		{token.FUNCTION, "fn"},
		// 		{token.LPARAN, "("},
		// 		{token.IDENT, "x"},
		// 		{token.COMMA, ","},
		// 		{token.IDENT, "y"},
		// 		{token.RPARAN, ")"},
		// 		{token.LBRACE, "{"},
		// 		{token.IDENT, "x"},
		// 		{token.PLUS, "+"},
		// 		{token.IDENT, "y"},
		// 		{token.SEMICOLON, ";"},
		// 		{token.RBRACE, "}"},
		// 		{token.SEMICOLON, ";"},
		// 		{token.LET, "let"},
		// 		{token.IDENT, "result"},
		// 		{token.ASSIGN, "="},
		// 		{token.IDENT, "add"},
		// 		{token.LPARAN, "("},
		// 		{token.IDENT, "five"},
		// 		{token.COMMA, ","},
		// 		{token.IDENT, "ten"},
		// 		{token.RPARAN, ")"},
		// 		{token.SEMICOLON, ";"},
		// 		{token.EOF, ""},
		// 	},
		// },
	}

	for i, tt := range tests {
		l := New(tt.input)
		var tok token.Token
		for j, want := range tt.wants {
			l, tok = l.NextToken()
			if tok.Type != want.ttype {
				t.Fatalf(
					"tests[%d, %d] - tokentype wrong. expected=%q, got=%q",
					i, j,
					want.ttype,
					tok.Type,
				)
			}
			if tok.Literal != want.literal {
				t.Fatalf(
					"tests[%d, %d] - literal wrong. expected=%q, got=%q",
					i, j,
					want.literal,
					tok.Literal,
				)
			}
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

func TestReadIdentifier(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `let`,
			want:  `let`,
		},
		{
			input: `l_et`,
			want:  `l_et`,
		},
		{
			input: `l1et`,
			want:  `l`,
		},
		{
			input: `1et`,
			want:  ``,
		},
	}

	for i, tt := range tests {
		l := New(tt.input)
		if _, got := l.readIdentifier(); got != tt.want {
			t.Fatalf(
				"tests[%d] - readIdentifier wrong. expected=%q, got=%q",
				i,
				tt.want,
				got,
			)
		}
	}
}

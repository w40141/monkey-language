package lexer

import (
	"reflect"
	"testing"

	"github.com/w40141/monkey-language/golang/token"
)

func TestReadChar(t *testing.T) {
	tests := []struct {
		input Lexer
		want  Lexer
	}{
		{
			input: Lexer{
				input:        "=1",
				position:     0,
				readPosition: 1,
				nowChar:      '=',
			},
			want: Lexer{
				input:        "=1",
				position:     1,
				readPosition: 2,
				nowChar:      '1',
			},
		},
		{
			input: Lexer{
				input:        "=",
				position:     0,
				readPosition: 1,
				nowChar:      '=',
			},
			want: Lexer{
				input:        "=",
				position:     1,
				readPosition: 2,
				nowChar:      0,
			},
		},
	}

	for i, tt := range tests {
		tt.input.readChar()
		if !reflect.DeepEqual(tt.input, tt.want) {
			t.Fatalf(
				"tests[%d] - lexer wrong. expected=%+v, got=%+v",
				i,
				tt.want,
				tt.input,
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
			input: `
			"foobar"
			"foo bar"
			`,
			wants: []want{
				{token.STRING, "foobar"},
				{token.STRING, "foo bar"},
				{token.EOF, ""},
			},
		},
		{
			input: `#`,
			wants: []want{
				{token.ILLEGAL, "#"},
			},
		},
		{
			input: `=+-*/!,;(){}<>`,
			wants: []want{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.MINUS, "-"},
				{token.TIMES, "*"},
				{token.DIVIDE, "/"},
				{token.BANG, "!"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.LPARAN, "("},
				{token.RPARAN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.LT, "<"},
				{token.GT, ">"},
				{token.EOF, ""},
			},
		},
		{
			input: `
			let five = 5;
			let ten = 10;
			let add = fn(x, y) {
				x + y;
			};
			let result = add(five, ten);
			`,
			wants: []want{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPARAN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPARAN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPARAN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPARAN, ")"},
				{token.SEMICOLON, ";"},
			},
		},
		{
			input: `
			if (5 < 10) {
				return true;
			} else {
				return false;
			}
			`,
			wants: []want{
				{token.IF, "if"},
				{token.LPARAN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPARAN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		{
			input: `
			==
			!=
			`,
			wants: []want{
				{token.EQ, "=="},
				{token.NQ, "!="},
				{token.EOF, ""},
			},
		},
		{
			input: `[1, 2]`,
			wants: []want{
				{token.LBRACKET, "["},
				{token.INT, "1"},
				{token.COMMA, ","},
				{token.INT, "2"},
				{token.RBRACKET, "]"},
				{token.EOF, ""},
			},
		},
		{
			input: `{"foo": "bar"}`,
			wants: []want{
				{token.LBRACE, "{"},
				{token.STRING, "foo"},
				{token.COLON, ":"},
				{token.STRING, "bar"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
	}

	for i, tt := range tests {
		l := New(tt.input)
		for j, want := range tt.wants {
			tok := l.NextToken()
			t.Logf("tests[%d, %d] - tok=%+v", i, j, tok)
			if tok.Type != want.ttype {
				t.Fatalf(
					"tests[%d, %d] - type wrong. expected=%q, got=%q",
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
		if got := l.readIdentifier(); got != tt.want {
			t.Fatalf(
				"tests[%d] - readIdentifier wrong. expected=%q, got=%q",
				i,
				tt.want,
				got,
			)
		}
	}
}

func TestSkipWhitespace(t *testing.T) {
	tests := []struct {
		input Lexer
		want  Lexer
	}{
		{
			input: Lexer{
				input:        "  =",
				position:     0,
				readPosition: 1,
				nowChar:      ' ',
			},
			want: Lexer{
				input:        "  =",
				position:     2,
				readPosition: 3,
				nowChar:      '=',
			},
		},
		{
			input: Lexer{
				input:        `	=`,
				position:     0,
				readPosition: 1,
				nowChar:      byte('\t'),
			},
			want: Lexer{
				input:        `	=`,
				position:     1,
				readPosition: 2,
				nowChar:      '=',
			},
		},
		{
			input: Lexer{
				input: `
=`,
				position:     0,
				readPosition: 1,
				nowChar:      byte('\n'),
			},
			want: Lexer{
				input: `
=`,
				position:     1,
				readPosition: 2,
				nowChar:      '=',
			},
		},
		{
			input: Lexer{
				input:        string(byte('\r')) + `=`,
				position:     0,
				readPosition: 1,
				nowChar:      byte('\r'),
			},
			want: Lexer{
				input:        string(byte('\r')) + `=`,
				position:     1,
				readPosition: 2,
				nowChar:      '=',
			},
		},
	}

	for i, tt := range tests {
		tt.input.skipWhitespace()
		if !reflect.DeepEqual(tt.input, tt.want) {
			t.Fatalf(
				"tests[%d] - lexer wrong. expected=%+v, got=%+v",
				i,
				tt.want,
				tt.input,
			)
		}
	}
}

func TestPeekChar(t *testing.T) {
	tests := []struct {
		input Lexer
		want  byte
	}{
		{
			input: Lexer{
				input:        "abc",
				position:     0,
				readPosition: 1,
				nowChar:      'a',
			},
			want: 'b',
		},
		{
			input: Lexer{
				input:        "abc",
				position:     1,
				readPosition: 2,
				nowChar:      'b',
			},
			want: 'c',
		},
		{
			input: Lexer{
				input:        "abc",
				position:     2,
				readPosition: 3,
				nowChar:      'c',
			},
			want: 0,
		},
	}
	for i, tt := range tests {
		if got := tt.input.peekChar(); got != tt.want {
			t.Fatalf(
				"tests[%d] - peekChar wrong. expected=%q, got=%q",
				i,
				tt.want,
				got,
			)
		}
	}
}

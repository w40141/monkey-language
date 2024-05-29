package parser

import (
	"testing"

	"github.com/w40141/monkey-language/golang/ast"
	"github.com/w40141/monkey-language/golang/lexer"
)

func TestLetStatements(t *testing.T) {
	type want struct {
		length      int
		identifiers []string
	}
	tests := []struct {
		name  string
		input string
		want  want
	}{
		{
			name:  "Single let statement",
			input: "let x = 5;",
			want: want{
				length:      1,
				identifiers: []string{"x"},
			},
		},
		{
			name: "Multi let statements",
			input: `
			let x = 5;
			let y = 10;
			let foobar = 838383;
			`,
			want: want{
				length: 3,
				identifiers: []string{
					"x",
					"y",
					"foobar",
				},
			},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		prg := p.ParseProgram()
		checkParserErrors(t, p)
		if prg == nil {
			t.Fatalf("ParseProgram() returned nil")
		}
		if l := len(prg.Statements); l != tt.want.length {
			t.Fatalf("len(prg.Statements) is not %d. got=%d", tt.want.length, l)
		}
		for i, stmt := range prg.Statements {
			name := tt.want.identifiers[i]
			if !testLetStatements(t, stmt, name) {
				return
			}
		}
	}
}

func testLetStatements(t *testing.T, stmt ast.Statement, name string) bool {
	if tl := stmt.TokenLiteral(); tl != "let" {
		t.Errorf("s.TokenLiteral() not 'let'. got=%q", tl)
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if tl := letStmt.Name.TokenLiteral(); tl != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %s. got=%s", name, tl)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	type want struct {
		length int
	}
	tests := []struct {
		name  string
		input string
		want  want
	}{
		{
			name:  "Single return statement",
			input: "return 5;",
			want: want{
				length: 1,
			},
		},
		{
			name: "Multi return statements",
			input: `
			return 5;
			return 10;
			return add(10);
			`,
			want: want{
				length: 3,
			},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		prg := p.ParseProgram()
		checkParserErrors(t, p)
		if prg == nil {
			t.Fatalf("ParseProgram() returned nil")
		}
		if l := len(prg.Statements); l != tt.want.length {
			t.Fatalf("len(prg.Statements) is not %d. got=%d", tt.want.length, l)
		}
		for _, stmt := range prg.Statements {
			returnStmt, ok := stmt.(*ast.ReturnStatement)
			if !ok {
				t.Errorf("s not *ast.ReturnStatement. got=%T", stmt)
			}
			if tl := returnStmt.TokenLiteral(); tl != "return" {
				t.Errorf("returnStmt.TokenLiteral() not 'return'. got=%s", tl)
			}
		}
	}
}

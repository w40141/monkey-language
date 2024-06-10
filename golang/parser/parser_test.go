package parser

import (
	"fmt"
	"testing"

	"github.com/w40141/monkey-language/golang/ast"
	"github.com/w40141/monkey-language/golang/lexer"
)

func TestIfExpression(t *testing.T) {
	t.Errorf("Not implemented")
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedBoolean bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		boolean, ok := stmt.Expression.(*ast.Boolean)
		if !ok {
			t.Fatalf("exp not *ast.Boolean. got=%T", stmt.Expression)
		}
		if boolean.Value != tt.expectedBoolean {
			t.Errorf("boolean.Value not %t. got=%t", tt.expectedBoolean,
				boolean.Value)
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"-a * b;",
			"((-a) * b)",
		},
		{
			"!-a;",
			"(!(-a))",
		},
		{
			"a + b + c;",
			"((a + b) + c)",
		},
		{
			"a + b - c;",
			"((a + b) - c)",
		},
		{
			"a * b * c;",
			"((a * b) * c)",
		},
		{
			"a * b / c;",
			"((a * b) / c)",
		},
		{
			"a + b / c;",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 > 4 != 3 < 4",
			"((5 > 4) != (3 < 4))",
		},
		{
			"3 + 4 * 5 ==  3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		prg := p.ParseProgram()
		checkParserErrors(t, p)

		if prg.String() != tt.want {
			t.Errorf("prg.String() is not %q. got=%q", tt.want, prg.String())
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	type want struct {
		length   int
		operator string
		left     any
		right    any
	}

	tests := []struct {
		input string
		want  want
	}{
		{
			input: "1 + 5;",
			want: want{
				length:   1,
				operator: "+",
				left:     1,
				right:    5,
			},
		},
		{
			input: "1 - 5;",
			want: want{
				length:   1,
				operator: "-",
				left:     1,
				right:    5,
			},
		},
		{
			input: "1 * 5;",
			want: want{
				length:   1,
				operator: "*",
				left:     1,
				right:    5,
			},
		},
		{
			input: "1 / 5;",
			want: want{
				length:   1,
				operator: "/",
				left:     1,
				right:    5,
			},
		},
		{
			input: "1 > 5;",
			want: want{
				length:   1,
				operator: ">",
				left:     1,
				right:    5,
			},
		},
		{
			input: "1 < 5;",
			want: want{
				length:   1,
				operator: "<",
				left:     1,
				right:    5,
			},
		},
		{
			input: "1 == 5;",
			want: want{
				length:   1,
				operator: "==",
				left:     1,
				right:    5,
			},
		},
		{
			input: "1 != 5;",
			want: want{
				length:   1,
				operator: "!=",
				left:     1,
				right:    5,
			},
		},
		{
			input: "true == true;",
			want: want{
				length:   1,
				operator: "==",
				left:     true,
				right:    true,
			},
		},
		{
			input: "true != false;",
			want: want{
				length:   1,
				operator: "!=",
				left:     true,
				right:    false,
			},
		},
		{
			input: "false == false;",
			want: want{
				length:   1,
				operator: "==",
				left:     false,
				right:    false,
			},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		prg := p.ParseProgram()
		checkParserErrors(t, p)

		if l := len(prg.Statements); l != tt.want.length {
			t.Fatalf("len(prg.Statements) is not %d. got=%d", tt.want.length, l)
		}

		stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("s not *ast.ExpressionStatement. got=%T", prg.Statements[0])
		}

		if !testInfixExpression(t, stmt.Expression, tt.want.left, tt.want.operator, tt.want.right) {
			return
		}
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	type want struct {
		length   int
		operator string
		value    any
	}

	tests := []struct {
		input string
		want  want
	}{
		{
			input: "-5;",
			want: want{
				length:   1,
				operator: "-",
				value:    5,
			},
		},
		{
			input: "!5;",
			want: want{
				length:   1,
				operator: "!",
				value:    5,
			},
		},
		{
			input: "!true;",
			want: want{
				length:   1,
				operator: "!",
				value:    true,
			},
		},
		{
			input: "!false;",
			want: want{
				length:   1,
				operator: "!",
				value:    false,
			},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		prg := p.ParseProgram()
		checkParserErrors(t, p)

		if l := len(prg.Statements); l != tt.want.length {
			t.Fatalf("len(prg.Statements) is not %d. got=%d", tt.want.length, l)
		}

		stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("s not *ast.ExpressionStatement. got=%T", prg.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("exp not *ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.want.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.want.operator, exp.Operator)
		}

		if !testLiteralExpressin(t, exp.Right, tt.want.value) {
			return
		}
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	type want struct {
		length  int
		value   int64
		literal string
	}
	tests := []struct {
		input string
		want  want
	}{
		{
			input: "5;",
			want: want{
				length:  1,
				value:   5,
				literal: "5",
			},
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		prg := p.ParseProgram()
		checkParserErrors(t, p)
		if l := len(prg.Statements); l != tt.want.length {
			t.Fatalf("len(prg.Statements) is not %d. got=%d", tt.want.length, l)
		}

		stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("s not *ast.ExpressionStatement. got=%T", prg.Statements[0])
		}

		literal, ok := stmt.Expression.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
		}
		if literal.Value != tt.want.value {
			t.Errorf("literal.Value not %d. got=%d", tt.want.value, literal.Value)
		}
		if tl := literal.TokenLiteral(); tl != tt.want.literal {
			t.Errorf("idnt.TokenLiteral() not %s. got=%s", tt.want.literal, tl)
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	type want struct {
		length int
		value  string
	}
	tests := []struct {
		input string
		want  want
	}{
		{
			input: "foobar;",
			want: want{
				length: 1,
				value:  "foobar",
			},
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		prg := p.ParseProgram()
		checkParserErrors(t, p)
		if l := len(prg.Statements); l != tt.want.length {
			t.Fatalf("len(prg.Statements) is not %d. got=%d", tt.want.length, l)
		}

		stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("s not *ast.ExpressionStatement. got=%T", prg.Statements[0])
		}

		idnt, ok := stmt.Expression.(*ast.Identifier)
		if !ok {
			t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
		}
		if idnt.Value != tt.want.value {
			t.Errorf("idnt.Value not %s. got=%s", tt.want.value, idnt.Value)
		}
		if tl := idnt.TokenLiteral(); tl != tt.want.value {
			t.Errorf("idnt.TokenLiteral() not %s. got=%s", tt.want.value, tl)
		}
	}
}

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

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral() not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
}

func testIdentifire(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.literal.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral() not %s. got=%s", value, ident.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpressin(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifire(t, exp, v)
	case bool:
		return testBoolean(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpressin(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpressin(t, opExp.Right, right) {
		return false
	}

	return true
}

func testBoolean(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.Token.Literal not %t. got=%s", value, bo.TokenLiteral())
		return false
	}

	return true
}

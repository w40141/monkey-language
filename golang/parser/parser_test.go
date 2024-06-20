package parser

import (
	"fmt"
	"testing"

	"github.com/w40141/monkey-language/golang/ast"
	"github.com/w40141/monkey-language/golang/lexer"
)

func TestParsingIndexExpression(t *testing.T) {
	input := "myArray[1 + 1]"
	prg := parseProgram(t, input)

	stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", prg.Statements[0])
	}

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
	}
	if !testIdentifire(t, indexExp.Left, "myArray") {
		return
	}

	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
		return
	}
}

func TestParsingArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	prg := parseProgram(t, input)

	stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", prg.Statements[0])
	}
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("exp not *ast.ArrayLiteral. got=%T", stmt.Expression)
	}
	if length := len(array.Elements); length != 3 {
		t.Fatalf("len(array.Elements) not 3. got=%d", length)
	}
	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world";`
	want := "hello world"
	prg := parseProgram(t, input)
	stmt := prg.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != want {
		t.Errorf("literal.Value not %q. got=%q", want, literal.Value)
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"

	prg := parseProgram(t, input)
	if length := len(prg.Statements); length != 1 {
		t.Fatalf("program.Statemenns does not contain %d statements. got=%d", 1, length)
	}

	stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			prg.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testIdentifire(t, exp.Function, "add") {
		return
	}

	if length := len(exp.Arguments); length != 3 {
		t.Fatalf("wrong length of arguments. got=%d", length)
	}

	testLiteralExpressin(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestFunctionParmeterParsing(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{input: "fn() {};", want: []string{}},
		{input: "fn(x) {};", want: []string{"x"}},
		{input: "fn(x, y, z) {};", want: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		prg := parseProgram(t, tt.input)
		stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				prg.Statements[0])
		}

		function, ok := stmt.Expression.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T", stmt.Expression)
		}

		if length := len(function.Parameters); length != len(tt.want) {
			t.Fatalf("function literal parameters wrong. want %d. got=%d", len(tt.want), length)
		}

		for i, ident := range tt.want {
			testLiteralExpressin(t, function.Parameters[i], ident)
		}
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`
	prg := parseProgram(t, input)
	if length := len(prg.Statements); length != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, length)
	}

	stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			prg.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T", stmt.Expression)
	}

	if length := len(function.Parameters); length != 2 {
		t.Fatalf("function literal parameters wrong. want 2. got=%d", length)
	}

	testLiteralExpressin(t, function.Parameters[0], "x")
	testLiteralExpressin(t, function.Parameters[1], "y")

	if length := len(function.Body.Statements); length != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d", length)
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T",
			function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestIfElseExpression(t *testing.T) {
	input := "if (x < y) { x } else { y }"
	prg := parseProgram(t, input)
	if length := len(prg.Statements); length != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, length)
	}

	stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			prg.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if length := len(exp.Consequence.Statements); length != 1 {
		t.Fatalf("consequence is not 1 statements. got=%d", length)
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf(
			"Statements[0] is not ast.ExpressionStatement. got=%T",
			exp.Consequence.Statements[0],
		)
	}

	if !testIdentifire(t, consequence.Expression, "x") {
		return
	}

	if length := len(exp.ALternative.Statements); length != 1 {
		t.Fatalf("consequence is not 1 statements. got=%d", length)
	}

	alternative, ok := exp.ALternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf(
			"Statements[0] is not ast.ExpressionStatement. got=%T",
			exp.ALternative.Statements[0],
		)
	}
	if !testIdentifire(t, alternative.Expression, "y") {
		return
	}
}

func TestIfExpression(t *testing.T) {
	input := "if (x < y) { x }"
	prg := parseProgram(t, input)
	if length := len(prg.Statements); length != 1 {
		t.Fatalf("program.Statemenns does not contain %d statements. got=%d", 1, length)
	}

	stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			prg.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if length := len(exp.Consequence.Statements); length != 1 {
		t.Fatalf("consequence is not 1 statements. got=%d", length)
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf(
			"Statements[0] is not ast.ExpressionStatement. got=%T",
			exp.Consequence.Statements[0],
		)
	}

	if !testIdentifire(t, consequence.Expression, "x") {
		return
	}

	if exp.ALternative != nil {
		t.Errorf("exp.ALternative.Statements was not nil. got=%+v",
			exp.ALternative)
	}
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
		prg := parseProgram(t, tt.input)
		if len(prg.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d",
				len(prg.Statements))
		}

		stmt, ok := prg.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				prg.Statements[0])
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
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
	}

	for _, tt := range tests {
		prg := parseProgram(t, tt.input)
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
		prg := parseProgram(t, tt.input)
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
		prg := parseProgram(t, tt.input)
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
		prg := parseProgram(t, tt.input)
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
		prg := parseProgram(t, tt.input)
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
	type identifier struct {
		idnt  string
		value any
	}
	type want struct {
		length      int
		identifiers []identifier
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
				length: 1,
				identifiers: []identifier{
					{"x", 5},
				},
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
				identifiers: []identifier{
					{"x", 5},
					{"y", 10},
					{"foobar", 838383},
				},
			},
		},
	}

	for _, tt := range tests {
		prg := parseProgram(t, tt.input)
		if prg == nil {
			t.Fatalf("ParseProgram() returned nil")
		}
		if l := len(prg.Statements); l != tt.want.length {
			t.Fatalf("len(prg.Statements) is not %d. got=%d", tt.want.length, l)
		}
		for i, stmt := range prg.Statements {
			identifier := tt.want.identifiers[i]
			if !testLetStatements(t, stmt, identifier.idnt) {
				return
			}

			val := stmt.(*ast.LetStatement).Value
			if !testLiteralExpressin(t, val, identifier.value) {
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

func parseProgram(t *testing.T, input string) *ast.Program {
	l := lexer.New(input)
	p := New(l)
	prg := p.ParseProgram()
	checkParserErrors(t, p)
	return prg
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

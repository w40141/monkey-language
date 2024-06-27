package evaluator

import (
	"testing"

	"github.com/w40141/monkey-language/golang/lexer"
	"github.com/w40141/monkey-language/golang/object"
	"github.com/w40141/monkey-language/golang/parser"
)

func TestHashIndexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{`{"foo": 5}["foo"]`, 5},
		{`{"foo": 5}["nil"]`, nil},
		{`let key = "foo"; {"foo": 5}[key]`, 5},
		{`{}["foo"]`, nil},
		{`{5: 5}[5]`, 5},
		{`{true: 5}[true]`, 5},
		{`{false: 5}[false]`, 5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestHashLiterals(t *testing.T) {
	input := `let two = "two";
	{
	"one": 10 - 9,
	two: 1 + 1,
	"thr" + "ee": 6 / 2,
	4: 4,
	true: 5,
	false: 6,
}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		(&object.Boolean{Value: true}).HashKey():   5,
		(&object.Boolean{Value: false}).HashKey():  6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	t.Log(result)
	for key, value := range expected {
		pair, ok := result.Pairs[key]
		if !ok {
			t.Fatalf("no pair for given key in Pairs. key=%+v, value=%+v", key, value)
		}
		testIntegerObject(t, pair.Value, value)
	}
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{input: "[1, 2, 3][0]", expected: 1},
		{input: "[1, 2, 3][1]", expected: 2},
		{input: "[1, 2, 3][2]", expected: 3},
		{input: "let i = 0; [1][i];", expected: 1},
		{input: "[1, 2, 3][1 + 1]", expected: 3},
		{input: "let myArray = [1, 2, 3]; myArray[2];", expected: 3},
		{input: "let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];", expected: 6},
		{input: "let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]", expected: 2},
		{input: "[1, 2, 3][3]", expected: nil},
		{input: "[1, 2, 3][-1]", expected: nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}
	if len(result.Elems) != 3 {
		t.Fatalf("array has wrong number of elements. got=%d", len(result.Elems))
	}
	testIntegerObject(t, result.Elems[0], 1)
	testIntegerObject(t, result.Elems[1], 4)
	testIntegerObject(t, result.Elems[2], 6)
}

func TestBuildinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{input: `len("")`, expected: 0},
		{input: `len("four")`, expected: 4},
		{input: `len("hello world")`, expected: 11},
		{input: `len(1)`, expected: "argument to `len` not supported, got INTEGER"},
		{input: `len("one", "two"),`, expected: "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			err, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%(+v))", evaluated, evaluated)
			}
			if err.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, err.Message)
			}
		}
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestClosures(t *testing.T) {
	input := `
let newAdder = fn(x) {
	fn(y) { x + y };
};
let addTwo = newAdder(2);
addTwo(2);`
	testIntegerObject(t, testEval(input), 4)
}

func TestFuntionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x){ 2 * x }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5);", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	type want struct {
		length int
		body   string
		params []string
	}
	tests := []struct {
		input string
		want  want
	}{
		{
			input: "fn(x) { x + 2; };",
			want: want{
				1,
				"(x + 2)",
				[]string{"x"},
			},
		},
		{
			input: "fn(x, y) { x + y + 2; };",
			want: want{
				2,
				"((x + y) + 2)",
				[]string{"x", "y"},
			},
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fn, ok := evaluated.(*object.Function)
		if !ok {
			t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
		}

		if len(fn.Parameters) != tt.want.length {
			t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
		}

		for i, param := range tt.want.params {
			if fn.Parameters[i].String() != param {
				t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[i])
			}
		}

		if fn.Body.String() != tt.want.body {
			t.Fatalf("body is not %s. got=%q", tt.want.body, fn.Body.String())
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; a", 5},
		{"let a = 5; let b = a; let c = a + b; c", 10},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5", 10},
		{"-5 + 10", 5},
		{"1 + 5 * 5", 26},
		{"1 * 5 + 5", 10},
		{"5 / 5 + 10", 11},
		{"10 + 5 / 5", 11},
		{"5 * 2 / 5", 2},
		{"2 * (5 + 5)", 20},
		{"4 * 5 / 2", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 == 2", false},
		{"1 != 1", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != true", false},
		{"true != false", true},
		{"false != false", false},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 11; 9;", 11},
		{"return 2 * 6; 9;", 12},
		{"9; return 13; 9;", 13},
		{"if (10 > 1) { if (10 > 1) { return 10; } return 1; }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHanding(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + true;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + true; 5", "type mismatch: INTEGER + BOOLEAN"},
		{"-true;", "unknown operator: -BOOLEAN"},
		{"true + false;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; true + false; 5", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if (10 > 1) { true + false; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "identifier not found: foobar"},
		{`"foobar" - "sss"`, "unknown operator: STRING - STRING"},
		{`{"name": "Monkey"}[fn(x) { x }];`, "unusable as hash key: FUNCTION"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expected {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expected, errObj.Message)
		}
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	pgm := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(pgm, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != nullObj {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

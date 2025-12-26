package evaluator

import (
	"lexicon/src/lexer"
	"lexicon/src/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"10 % 3", 1},
		{"2 ** 3", 8},
		{"2 ** 3 ** 2", 512},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.5", 5.5},
		{"10.2", 10.2},
		{"-5.5", -5.5},
		{"5.5 + 2.5", 8.0},
		{"5.5 - 2.5", 3.0},
		{"5.5 * 2.0", 11.0},
		{"10.0 / 2.0", 5.0},
		{"2.0 ** 3.0", 8.0},
		{"5 + 2.5", 7.5},
		{"10.5 - 5", 5.5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
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
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"5 <= 5", true},
		{"5 >= 5", true},
		{"5 <= 4", false},
		{"5 >= 6", false},
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

func TestLogicalOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true && true", true},
		{"true && false", false},
		{"false && true", false},
		{"false && false", false},
		{"true || true", true},
		{"true || false", true},
		{"false || true", true},
		{"false || false", false},
		{"(1 < 2) && (2 < 3)", true},
		{"(1 > 2) || (2 < 3)", true},
		{"(1 > 2) && (2 < 3)", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { sprout x = 10; }", 10},
		{"if (false) { sprout x = 10; }", nil},
		{"if (1) { sprout x = 10; }", 10},
		{"if (1 < 2) { sprout x = 10; }", 10},
		{"if (1 > 2) { sprout x = 10; }", nil},
		{"if (1 > 2) { sprout x = 10; } else { sprout y = 20; }", 20},
		{"if (1 < 2) { sprout x = 10; } else { sprout y = 20; }", 10},
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

func TestVariableDeclarations(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"sprout a = 5; a;", 5},
		{"sprout a = 5 * 5; a;", 25},
		{"sprout a = 5; sprout b = a; b;", 5},
		{"sprout a = 5; sprout b = a; sprout c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestVariableReassignment(t *testing.T) {
	input := `
	sprout x = 10;
	x = 20;
	x;
	`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 20)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`
	evaluated := testEval(input)

	str, ok := evaluated.(*String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	evaluated := testEval(input)

	str, ok := evaluated.(*String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			"10 / 0",
			"division by zero",
		},
		{
			"10 % 0",
			"modulo by zero",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestNestedScopes(t *testing.T) {
	input := `
	sprout x = 10;
	if (true) {
		sprout y = 20;
		x = x + y;
	}
	x;
	`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 30)
}

func TestComplexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`sprout x = 5;
			sprout y = 10;
			if (x < y) {
				sprout z = x + y;
				z;
			}`,
			15,
		},
		{
			`sprout a = 5;
			sprout b = 10;
			sprout c = 15;
			if (a < b && b < c) {
				a + b + c;
			}`,
			30,
		},
		{
			`sprout x = 10;
			if (x > 5) {
				sprout y = x * 2;
				if (y > 15) {
					y + 5;
				}
			}`,
			25,
		},
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

// helper functions
func testEval(input string) Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj Object, expected int64) bool {
	result, ok := obj.(*Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func testFloatObject(t *testing.T, obj Object, expected float64) bool {
	result, ok := obj.(*Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f",
			result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj Object, expected bool) bool {
	result, ok := obj.(*Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

package parser

import (
	"lexicon/src/ast"
	"lexicon/src/lexer"
	"testing"
)

func TestVariableDeclaration(t *testing.T) {
	input := `
	sprout x = 10;
	sprout y int = 20;
	x = 30;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		for i, stmt := range program.Statements {
			t.Logf("Statement %d: %T - %v", i, stmt, stmt)
		}
		t.Fatalf("expected 3 statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got=%T", program.Statements[0])
	}
	if stmt.Name.Value != "x" {
		t.Errorf("expected variable name 'x', got=%s", stmt.Name.Value)
	}
}

func TestIfElseParsing(t *testing.T) {
	input := `
	if (x) {
		sprout y = 10;
	} else {
		sprout z = 20;
	}
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got=%d", len(program.Statements))
	}

	ifStmt, ok := program.Statements[0].(*ast.IfExpression)
	if !ok {
		t.Fatalf("expected *ast.IfExpression, got=%T", program.Statements[0])
	}

	ident, ok := ifStmt.Condition.(*ast.Identifier)
	if !ok || ident.Value != "x" {
		t.Errorf("expected condition to be identifier 'x', got=%v", ifStmt.Condition)
	}

	if len(ifStmt.Consequence.Statements) != 1 {
		t.Errorf("expected 1 statement in consequence, got=%d", len(ifStmt.Consequence.Statements))
	}

	if len(ifStmt.Alternative.Statements) != 1 {
		t.Errorf("expected 1 statement in alternative, got=%d", len(ifStmt.Alternative.Statements))
	}
}

// Test Print Statements
func TestPrintStatements(t *testing.T) {
	input := `
	echo "Hello, World!";
	echo 42;
	echo x + y;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.PrintStatement)
	if !ok {
		t.Fatalf("expected *ast.PrintStatement, got=%T", program.Statements[0])
	}

	literal, ok := stmt.Value.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("expected string literal, got=%T", stmt.Value)
	}
	if literal.Value != "Hello, World!" {
		t.Errorf("expected 'Hello, World!', got=%s", literal.Value)
	}
}

func TestSimpleExpressions(t *testing.T) {
	input := `
	sprout x = 5 + 3;
	sprout y = 10 * 2;
	sprout z = x + y;
	a = 3**2;
	a =-y;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 5 {
		t.Fatalf("expected 3 statements, got=%d", len(program.Statements))
	}

	// Test addition
	firstStmt, ok := program.Statements[0].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got=%T", program.Statements[0])
	}

	infixExpr, ok := firstStmt.Value.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expected infix expression, got=%T", firstStmt.Value)
	}

	if infixExpr.Operator != "+" {
		t.Errorf("expected '+' operator, got=%s", infixExpr.Operator)
	}
}

func TestLogicalExpressions(t *testing.T) {
	input := `
	sprout x = true && false;
	sprout y = true || false;
	z = x and y;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("expected 2 statements, got=%d", len(program.Statements))
	}

	// Test logical AND
	firstStmt, ok := program.Statements[0].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got=%T", program.Statements[0])
	}

	infixExpr, ok := firstStmt.Value.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expected infix expression, got=%T", firstStmt.Value)
	}

	if infixExpr.Operator != "&&" {
		t.Errorf("expected '&&' operator, got=%s", infixExpr.Operator)
	}

	// Test logical OR
	secondStmt, ok := program.Statements[1].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got=%T", program.Statements[1])
	}

	orInfixExpr, ok := secondStmt.Value.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expected infix expression, got=%T", secondStmt.Value)
	}

	if orInfixExpr.Operator != "||" {
		t.Errorf("expected '||' operator, got=%s", orInfixExpr.Operator)
	}
}

func TestGroupedExpression(t *testing.T) {
	input := `
	z = (x + y);
	z = (x and y);
	z = ((x + y) * 2);
	z = (x + y) and (a * b);
	a = (b + c) * d;
	a = (b + c) * d;
	z = x * (y + 2);
	result = (a + b) and (c * d);
	q = not x;
	t = !x;
	b = not (x + y);
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 11 {
		t.Fatalf("expected 11 statements, got=%d", len(program.Statements))
	}

	// helper for infix
	checkInfixExpression := func(stmt ast.Statement, expected string) {
		decl, ok := stmt.(*ast.VariableDeclaration)
		if !ok {
			t.Fatalf("expected *ast.VariableDeclaration, got=%T", stmt)
		}
		if _, ok := decl.Value.(*ast.InfixExpression); !ok {
			t.Errorf("expected infix expression for %s, got=%T", expected, decl.Value)
		}
	}

	// helper for prefix
	checkPrefixExpression := func(stmt ast.Statement, expected string) {
		decl, ok := stmt.(*ast.VariableDeclaration)
		if !ok {
			t.Fatalf("expected *ast.VariableDeclaration, got=%T", stmt)
		}
		if _, ok := decl.Value.(*ast.PrefixExpression); !ok {
			t.Errorf("expected prefix expression for %s, got=%T", expected, decl.Value)
		}
	}

	checkInfixExpression(program.Statements[0], "(x + y)")
	checkInfixExpression(program.Statements[1], "(x and y)")
	checkInfixExpression(program.Statements[2], "((x + y) * 2)")
	checkInfixExpression(program.Statements[3], "(x + y) and (a * b)")
	checkInfixExpression(program.Statements[4], "(b + c) * d")
	checkInfixExpression(program.Statements[5], "(b + c) * d")
	checkInfixExpression(program.Statements[6], "x * (y + 2)")
	checkInfixExpression(program.Statements[7], "(a + b) and (c * d)")
	checkPrefixExpression(program.Statements[8], "not x")
	checkPrefixExpression(program.Statements[9], "!x")
	checkPrefixExpression(program.Statements[10], "not (x + y)")
}

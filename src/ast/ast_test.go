package ast

import (
	"lexicon/src/token"
	"testing"
)

func TestProgramString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VariableDeclaration{
				Token: token.Token{Type: token.SPROUT, Literal: "sprout"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
				Type:  nil,
				Value: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "10"}, Value: 10},
			},
		},
	}

	expected := "sprout x = 10;"
	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q, want=%q", program.String(), expected)
	}
}

func TestProgramTokenLiteral(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VariableDeclaration{
				Token: token.Token{Type: token.SPROUT, Literal: "sprout"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
			},
		},
	}

	if program.TokenLiteral() != "sprout" {
		t.Errorf("program.TokenLiteral() = %q, want 'sprout'", program.TokenLiteral())
	}

	emptyProgram := &Program{Statements: []Statement{}}
	if emptyProgram.TokenLiteral() != "" {
		t.Errorf("empty program.TokenLiteral() = %q, want ''", emptyProgram.TokenLiteral())
	}
}

func TestVariableDeclarationString(t *testing.T) {
	tests := []struct {
		name     string
		vd       *VariableDeclaration
		expected string
	}{
		{
			name: "simple variable",
			vd: &VariableDeclaration{
				Token: token.Token{Type: token.SPROUT, Literal: "sprout"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
				Value: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
			},
			expected: "sprout x = 5;",
		},
		{
			name: "typed variable",
			vd: &VariableDeclaration{
				Token: token.Token{Type: token.SPROUT, Literal: "sprout"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "age"}, Value: "age"},
				Type:  &Identifier{Token: token.Token{Type: token.TYPE_IDENT, Literal: "int"}, Value: "int"},
				Value: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "25"}, Value: 25},
			},
			expected: "sprout age int = 25;",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.vd.String()
			if result != tt.expected {
				t.Errorf("VariableDeclaration.String() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestIdentifier(t *testing.T) {
	ident := &Identifier{
		Token: token.Token{Type: token.IDENT, Literal: "myVar"},
		Value: "myVar",
	}

	if ident.TokenLiteral() != "myVar" {
		t.Errorf("TokenLiteral() = %q, want 'myVar'", ident.TokenLiteral())
	}

	if ident.String() != "myVar" {
		t.Errorf("String() = %q, want 'myVar'", ident.String())
	}
}

func TestIntegerLiteral(t *testing.T) {
	intLit := &IntegerLiteral{
		Token: token.Token{Type: token.INT, Literal: "42"},
		Value: 42,
	}

	if intLit.TokenLiteral() != "42" {
		t.Errorf("TokenLiteral() = %q, want '42'", intLit.TokenLiteral())
	}

	if intLit.String() != "42" {
		t.Errorf("String() = %q, want '42'", intLit.String())
	}
}

func TestFloatLiteral(t *testing.T) {
	floatLit := &FloatLiteral{
		Token: token.Token{Type: token.FLOAT, Literal: "3.14"},
		Value: 3.14,
	}

	if floatLit.TokenLiteral() != "3.14" {
		t.Errorf("TokenLiteral() = %q, want '3.14'", floatLit.TokenLiteral())
	}

	if floatLit.String() != "3.140000" {
		t.Errorf("String() = %q, want '3.140000'", floatLit.String())
	}
}

func TestBooleanLiteral(t *testing.T) {
	tests := []struct {
		value    bool
		expected string
	}{
		{true, "true"},
		{false, "false"},
	}

	for _, tt := range tests {
		boolLit := &BooleanLiteral{
			Token: token.Token{Type: token.BOOL, Literal: tt.expected},
			Value: tt.value,
		}

		if boolLit.String() != tt.expected {
			t.Errorf("BooleanLiteral.String() = %q, want %q", boolLit.String(), tt.expected)
		}
	}
}

func TestStringLiteral(t *testing.T) {
	strLit := &StringLiteral{
		Token: token.Token{Type: token.STRING, Literal: "hello"},
		Value: "hello",
	}

	if strLit.TokenLiteral() != "hello" {
		t.Errorf("TokenLiteral() = %q, want 'hello'", strLit.TokenLiteral())
	}

	if strLit.String() != `"hello"` {
		t.Errorf("String() = %q, want %q", strLit.String(), `"hello"`)
	}
}

func TestPrefixExpression(t *testing.T) {
	prefixExpr := &PrefixExpression{
		Token:    token.Token{Type: token.MINUS, Literal: "-"},
		Operator: "-",
		Right:    &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
	}

	expected := "(-5)"
	if prefixExpr.String() != expected {
		t.Errorf("PrefixExpression.String() = %q, want %q", prefixExpr.String(), expected)
	}
}

func TestInfixExpression(t *testing.T) {
	infixExpr := &InfixExpression{
		Token:    token.Token{Type: token.PLUS, Literal: "+"},
		Left:     &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		Operator: "+",
		Right:    &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "3"}, Value: 3},
	}

	expected := "(5 + 3)"
	if infixExpr.String() != expected {
		t.Errorf("InfixExpression.String() = %q, want %q", infixExpr.String(), expected)
	}
}

func TestBlockStatement(t *testing.T) {
	block := &BlockStatement{
		Token: token.Token{Type: token.LBRACE, Literal: "{"},
		Statements: []Statement{
			&ExpressionStatement{
				Token:      token.Token{Type: token.INT, Literal: "5"},
				Expression: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
			},
		},
	}

	if block.TokenLiteral() != "{" {
		t.Errorf("BlockStatement.TokenLiteral() = %q, want '{'", block.TokenLiteral())
	}

	if len(block.Statements) != 1 {
		t.Errorf("BlockStatement has %d statements, want 1", len(block.Statements))
	}
}

func TestIfExpression(t *testing.T) {
	ifExpr := &IfExpression{
		Token: token.Token{Type: token.IF, Literal: "if"},
		Condition: &BooleanLiteral{
			Token: token.Token{Type: token.TRUE, Literal: "true"},
			Value: true,
		},
		Consequence: &BlockStatement{
			Token:      token.Token{Type: token.LBRACE, Literal: "{"},
			Statements: []Statement{},
		},
		Alternative: nil,
	}

	if ifExpr.TokenLiteral() != "if" {
		t.Errorf("IfExpression.TokenLiteral() = %q, want 'if'", ifExpr.TokenLiteral())
	}
}

func TestPrintStatement(t *testing.T) {
	printStmt := &PrintStatement{
		Token: token.Token{Type: token.ECHO, Literal: "echo"},
		Value: &StringLiteral{
			Token: token.Token{Type: token.STRING, Literal: "hello"},
			Value: "hello",
		},
	}

	expected := `echo "hello";`
	if printStmt.String() != expected {
		t.Errorf("PrintStatement.String() = %q, want %q", printStmt.String(), expected)
	}
}



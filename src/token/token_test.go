package token

import "testing"

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		{"echo", ECHO},
		{"sprout", SPROUT},
		{"if", IF},
		{"else", ELSE},
		{"and", LOGICAL_AND},
		{"or", LOGICAL_OR},
		{"not", LOGICAL_NOT},
		{"true", TRUE},
		{"false", FALSE},
		{"int", TYPE_IDENT},
		{"float", TYPE_IDENT},
		{"string", TYPE_IDENT},
		{"bool", TYPE_IDENT},
		{"x", IDENT},
		{"myVar", IDENT},
		{"variable123", IDENT},
	}

	for _, tt := range tests {
		result := LookupIdent(tt.input)
		if result != tt.expected {
			t.Errorf("LookupIdent(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestTokenType(t *testing.T) {
	// Test that token types are correctly defined
	tests := []struct {
		tokenType TokenType
		expected  string
	}{
		{ILLEGAL, "ILLEGAL"},
		{EOF, "EOF"},
		{IDENT, "IDENT"},
		{INT, "INT"},
		{STRING, "STRING"},
		{FLOAT, "FLOAT"},
		{ASSIGN, "="},
		{PLUS, "+"},
		{MINUS, "-"},
		{MUL, "*"},
		{DIV, "/"},
		{MOD, "%"},
		{EXP, "**"},
		{GT, ">"},
		{LT, "<"},
		{EQ, "=="},
		{NOT_EQ, "!="},
		{LTE, "<="},
		{GTE, ">="},
		{LOGICAL_AND, "&&"},
		{LOGICAL_OR, "||"},
		{LOGICAL_NOT, "!"},
		{TRUE, "TRUE"},
		{FALSE, "FALSE"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{ECHO, "ECHO"},
		{SPROUT, "SPROUT"},
		{IF, "IF"},
		{ELSE, "ELSE"},
		{COMMENT, "COMMENT"},
	}

	for _, tt := range tests {
		if string(tt.tokenType) != tt.expected {
			t.Errorf("TokenType %q != %q", tt.tokenType, tt.expected)
		}
	}
}

func TestToken(t *testing.T) {
	tok := Token{
		Type:    INT,
		Literal: "42",
		Line:    1,
		Column:  5,
	}

	if tok.Type != INT {
		t.Errorf("Token.Type = %q, want INT", tok.Type)
	}
	if tok.Literal != "42" {
		t.Errorf("Token.Literal = %q, want '42'", tok.Literal)
	}
	if tok.Line != 1 {
		t.Errorf("Token.Line = %d, want 1", tok.Line)
	}
	if tok.Column != 5 {
		t.Errorf("Token.Column = %d, want 5", tok.Column)
	}
}

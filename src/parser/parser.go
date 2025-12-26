package parser

import (
	"fmt"
	"lexicon/src/ast"
	"lexicon/src/lexer"
	"lexicon/src/token"
	"strconv"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors    []string // parsing errors with line numbers
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Errors returns the list of parsing errors
func (p *Parser) Errors() []string {
	return p.errors
}

// addError adds a formatted error message with line number
func (p *Parser) addError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	p.errors = append(p.errors, msg)
}

// peekError adds an error when expected token doesn't match
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("[Line %d:%d] Expected next token to be %s, got %s instead",
		p.peekToken.Line, p.peekToken.Column, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		// skip semicolons and continue (nextToken already called in skip)
		if p.currToken.Type == token.SEMICOLON {
			p.nextToken()
			continue
		}

		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.COMMENT:
		// skip comments - they don't produce statements
		return nil
	case token.IDENT:
		// check if it's an assignment or expression
		if p.peekToken.Type == token.ASSIGN {
			return p.parseAssignment()
		}
		return p.parseExpressionStatement()
	case token.SPROUT:
		return p.parseVariableDeclaration()
	case token.ECHO:
		return p.parsePrintStatement()
	case token.IF:
		return p.parseIfExpression()
	default:
		return p.parseExpressionStatement()
	}
}

// sprout x = 10 | sprout x int = 10
func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	stmt := &ast.VariableDeclaration{Token: p.currToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if p.peekToken.Type == token.TYPE_IDENT {
		p.nextToken()
		stmt.Type = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}

// x = 10
func (p *Parser) parseAssignment() *ast.VariableDeclaration {
	stmt := &ast.VariableDeclaration{Token: p.currToken}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}

// if (cond) { } else { }
func (p *Parser) parseIfExpression() *ast.IfExpression {
	expr := &ast.IfExpression{Token: p.currToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expr.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expr.Consequence = p.parseBlockStatement()

	if p.peekToken.Type == token.ELSE {
		p.nextToken()
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		expr.Alternative = p.parseBlockStatement()
	}

	return expr
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currToken}
	p.nextToken()

	for p.currToken.Type != token.RBRACE && p.currToken.Type != token.EOF {
		// skip semicolons
		if p.currToken.Type == token.SEMICOLON {
			p.nextToken()
			continue
		}

		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

// echo "Hello"
func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	stmt := &ast.PrintStatement{Token: p.currToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}

// expression statement (for standalone expressions like "a;" or "5;")
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}
	stmt.Expression = p.parseExpression(LOWEST)

	// if expression is nil, return nil (no valid expression to parse)
	if stmt.Expression == nil {
		return nil
	}

	return stmt
}

// Arithmetic & Logical Expressions
// precedence constants and mapping
const (
	_ int = iota
	LOWEST
	LOGICAL_OR  // or, ||
	LOGICAL_AND // and, &&
	EQUALITY    // ==, !=
	COMPARISON  // <, >, <=, >=
	SUM         // +, -
	PRODUCT     // *, /, %
	EXPONENT    // **
	PREFIX      // -X, not X
)

var precedences = map[token.TokenType]int{
	// Logical Operators
	token.LOGICAL_OR:  LOGICAL_OR,
	token.LOGICAL_AND: LOGICAL_AND,
	token.OR:          LOGICAL_OR,
	token.AND:         LOGICAL_AND,

	// Comparison Operators
	token.EQ:     EQUALITY,
	token.NOT_EQ: EQUALITY,
	token.LT:     COMPARISON,
	token.GT:     COMPARISON,
	token.LTE:    COMPARISON,
	token.GTE:    COMPARISON,

	// Arithmetic Operators
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.MUL:   PRODUCT,
	token.DIV:   PRODUCT,
	token.MOD:   PRODUCT,
	token.EXP:   EXPONENT,
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	var leftExp ast.Expression

	switch p.currToken.Type {
	case token.INT:
		leftExp = p.parseIntegerLiteral()
	case token.FLOAT:
		leftExp = p.parseFloatLiteral()
	case token.IDENT:
		leftExp = p.parseIdentifier()
	case token.MINUS, token.LOGICAL_NOT, token.NOT:
		leftExp = p.parsePrefixExpression()
	case token.STRING:
		leftExp = p.parseStringLiteral()
	case token.TRUE, token.FALSE:
		leftExp = p.parseBooleanLiteral()
	case token.LPAREN:
		leftExp = p.parseGroupedExpression()
	default:
		return nil
	}

	for precedence < p.peekPrecedence() {
		switch p.peekToken.Type {
		case token.PLUS, token.MINUS, token.MUL, token.DIV, token.MOD, token.EXP,
			token.EQ, token.NOT_EQ, token.LT, token.GT, token.LTE, token.GTE,
			token.LOGICAL_AND, token.LOGICAL_OR,
			token.AND, token.OR:
			p.nextToken()
			leftExp = p.parseInfixExpression(leftExp)
		default:
			return leftExp
		}
	}

	return leftExp
}

// handle expressions in parenthesis
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token:    p.currToken,
		Left:     left,
		Operator: p.normalizeLogicalOperator(p.currToken.Literal), // to handle multiple ways of writing logical operators
	}

	precedence := p.currPrecedence()
	p.nextToken()

	// for right-associative operators like exponentiation
	if expr.Operator == "**" {
		expr.Right = p.parseExpression(precedence - 1)
	} else if p.currToken.Type == token.LPAREN {
		// to parse parenthesis
		expr.Right = p.parseGroupedExpression()
	} else {
		expr.Right = p.parseExpression(precedence)
	}

	return expr
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.normalizeLogicalOperator(p.currToken.Literal), // logical not condition
	}

	p.nextToken()

	// to parse parenthesis
	if p.currToken.Type == token.LPAREN {
		expr.Right = p.parseGroupedExpression()
	} else {
		expr.Right = p.parseExpression(PREFIX)
	}

	return expr
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{
		Token: p.currToken,
		Value: p.currToken.Type == token.TRUE,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.ParseInt(p.currToken.Literal, 10, 64)
	if err != nil {
		p.addError("[Line %d:%d] Could not parse %q as integer",
			p.currToken.Line, p.currToken.Column, p.currToken.Literal)
		return nil
	}
	return &ast.IntegerLiteral{Token: p.currToken, Value: value}
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	value, err := strconv.ParseFloat(p.currToken.Literal, 64)
	if err != nil {
		p.addError("[Line %d:%d] Could not parse %q as float",
			p.currToken.Line, p.currToken.Column, p.currToken.Literal)
		return nil
	}
	return &ast.FloatLiteral{
		Token: p.currToken,
		Value: value,
	}
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) currPrecedence() int {
	if prec, ok := precedences[p.currToken.Type]; ok {
		return prec
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.peekToken.Type]; ok {
		return prec
	}
	return LOWEST
}

func (p *Parser) normalizeLogicalOperator(op string) string {
	switch op {
	case "and", "&&":
		return "&&"
	case "or", "||":
		return "||"
	case "not", "!":
		return "!"
	default:
		return op
	}
}

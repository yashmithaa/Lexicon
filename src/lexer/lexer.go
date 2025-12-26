package lexer

import (
	"lexicon/src/token"
	"strings"
	"unicode"
)

type Lexer struct {
	input   string // source code
	currPos int    // current charecter position
	nextPos int    // next charecter position
	ch      rune
	line    int // current line number
	column  int // current column number
}

func New(input string) *Lexer {
	l := new(Lexer)
	l.input = input
	l.line = 1
	l.column = 0
	l.readChar()
	return l
}

// Reads next character
func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = rune(l.input[l.nextPos])
	}

	// Update line and column tracking
	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}

	l.currPos = l.nextPos
	l.nextPos++
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.currPos
	for unicode.IsLetter(rune(l.ch)) {
		l.readChar()
	}
	return l.input[start:l.currPos]
}

func (l *Lexer) readNumber() string {
	start := l.currPos
	hasDot := false

	for unicode.IsDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if hasDot {
				return "ERROR: Multiple decimal points in number"
			}
			hasDot = true
		}
		l.readChar()
	}

	return l.input[start:l.currPos]
}

func (l *Lexer) readComment() string {
	start := l.currPos
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[start:l.currPos]
}

// for multi character tokens like !=, ==
func (l *Lexer) peekChar() rune {
	if l.nextPos >= len(l.input) {
		return 0 // EOF
	}
	return rune(l.input[l.nextPos])
}

// creates token.Token to reduce code duplication
func (l *Lexer) newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: ch,
		Line:    l.line,
		Column:  l.column,
	}
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '#':
		comment := l.readComment()
		tok = token.Token{Type: token.COMMENT, Literal: comment}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.ASSIGN, "=")
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.LOGICAL_NOT, "!")
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.LOGICAL_AND, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.ILLEGAL, "&")
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.LOGICAL_OR, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.ILLEGAL, "|")
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.GTE, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.GT, ">")
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.LTE, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.LT, "<")
		}
	case '+':
		tok = l.newToken(token.PLUS, "+")
	case '-':
		tok = l.newToken(token.MINUS, "-")
	case '*':
		if l.peekChar() == '*' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.EXP, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.MUL, "*")
		}
	case '/':
		tok = l.newToken(token.DIV, "/")
	case '%':
		tok = l.newToken(token.MOD, "%")
	case '(':
		tok = l.newToken(token.LPAREN, "(")
	case ')':
		tok = l.newToken(token.RPAREN, ")")
	case '{':
		tok = l.newToken(token.LBRACE, "{")
	case '}':
		tok = l.newToken(token.RBRACE, "}")
	case ',':
		tok = l.newToken(token.COMMA, ",")
	case ';':
		tok = l.newToken(token.SEMICOLON, ";")
	case ':':
		tok = l.newToken(token.COLON, ":")
	case '.':
		tok = l.newToken(token.DOT, ".")
	case '"':
		return l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(rune(l.ch)) {
			tokLine := l.line
			tokCol := l.column
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // Efficient keyword check
			tok.Line = tokLine
			tok.Column = tokCol

			// Handle boolean literals
			if tok.Literal == "true" {
				tok.Type = token.TRUE
			} else if tok.Literal == "false" {
				tok.Type = token.FALSE
			}

			return tok
		} else if unicode.IsDigit(rune(l.ch)) {
			tokLine := l.line
			tokCol := l.column
			tok.Literal = l.readNumber()
			tok.Line = tokLine
			tok.Column = tokCol
			// check if it's a float (contains a decimal point)
			if strings.Contains(tok.Literal, ".") {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readString() token.Token {
	tokLine := l.line
	tokCol := l.column
	l.readChar()

	var value strings.Builder
	escaped := false

	for l.ch != 0 {
		if escaped {
			switch l.ch {
			case 'n':
				value.WriteRune('\n')
			case 't':
				value.WriteRune('\t')
			case 'r':
				value.WriteRune('\r')
			case '"':
				value.WriteRune('"')
			case '\\':
				value.WriteRune('\\')
			default:
				// Unrecognized escape sequence, could be an error
				value.WriteRune('\\')
				value.WriteRune(l.ch)
			}
			escaped = false
		} else if l.ch == '\\' {
			escaped = true
		} else if l.ch == '"' {
			break
		} else {
			value.WriteRune(l.ch)
		}
		l.readChar()
	}

	if l.ch != '"' {
		// Unterminated string
		return token.Token{
			Type:    token.ILLEGAL,
			Literal: "Unterminated string",
			Line:    tokLine,
			Column:  tokCol,
		}
	}

	l.readChar() // consume closing quote
	return token.Token{
		Type:    token.STRING,
		Literal: value.String(),
		Line:    tokLine,
		Column:  tokCol,
	}
}

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Datatypes for Variable Declaration
	IDENT      = "IDENT"
	INT        = "INT"
	STRING     = "STRING"
	FLOAT      = "FLOAT"
	BOOL       = "BOOL"
	TYPE_IDENT = "TYPE_IDENT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MUL    = "*"
	DIV    = "/"
	MOD    = "%"
	EXP    = "**"

	// Comparison Operators
	GT     = ">"
	LT     = "<"
	EQ     = "=="
	NOT_EQ = "!="
	LTE    = "<="
	GTE    = ">="

	// Bitwise Operators
	// XOR = "^"
	// AND = "&"
	// OR  = "|"
	// NOT = "!"

	// Logical Operators
	LOGICAL_AND = "&&"
	LOGICAL_OR  = "||"
	LOGICAL_NOT = "!"

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"

	// Boolean Literals
	TRUE  = "TRUE"
	FALSE = "FALSE"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COLON     = ":"
	DOT       = "."

	// Keywords
	ECHO   = "ECHO"
	SPROUT = "SPROUT"
	IF     = "IF"
	ELSE   = "ELSE"

	COMMENT = "COMMENT"
)

var keywords = map[string]TokenType{
	"echo":   ECHO,
	"sprout": SPROUT,

	"if":   IF,
	"else": ELSE,

	"and": LOGICAL_AND, // alternate for &&
	"or":  LOGICAL_OR,  // alternate for ||
	"not": LOGICAL_NOT,

	"true":  TRUE,
	"false": FALSE,

	"int":    TYPE_IDENT,
	"float":  TYPE_IDENT,
	"string": TYPE_IDENT,
	"bool":   TYPE_IDENT,
}

// To check if the identifier is a keyword or not
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

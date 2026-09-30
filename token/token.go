// Package token defines the data structures and helpers used to parse tokens from source.
package token

// Type is a category of token recognised by the lexer.
type Type string

// Token is a lexeme combined with its type.
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL marks a character the lexer does not recognise.
	ILLEGAL = "ILLEGAL"
	// EOF marks the end of the input.
	EOF = "EOF"

	// IDENT is a user-defined name such as a variable or function (add, foobar, x, y, ...).
	IDENT = "IDENT"
	// INT is an integer literal.
	INT = "INT"

	// ASSIGN is the assignment operator.
	ASSIGN = "="
	// PLUS is the addition operator.
	PLUS = "+"
	// MINUS is the subtraction operator.
	MINUS = "-"
	// BANG is the logical-not operator.
	BANG = "!"
	// ASTERISK is the multiplication operator.
	ASTERISK = "*"
	// SLASH is the division operator.
	SLASH = "/"

	// EQ is the equality comparison operator.
	EQ = "=="
	// NEQ is the inequality comparison operator.
	NEQ = "!="
	// LT is the less-than comparison operator.
	LT = "<"
	// GT is the greater-than comparison operator.
	GT = ">"

	// COMMA separates items in a list.
	COMMA = ","
	// SEMICOLON terminates a statement.
	SEMICOLON = ";"

	// LPAREN opens a parameter or argument list.
	LPAREN = "("
	// RPAREN closes a parameter or argument list.
	RPAREN = ")"
	// LBRACE opens a block.
	LBRACE = "{"
	// RBRACE closes a block.
	RBRACE = "}"

	// FUNCTION is the fn keyword.
	FUNCTION = "FUNCTION"
	// LET is the let keyword.
	LET = "LET"
	// TRUE is the true keyword.
	TRUE = "TRUE"
	// FALSE is the false keyword.
	FALSE = "FALSE"
	// IF is the if keyword.
	IF = "IF"
	// ELSE is the else keyword.
	ELSE = "ELSE"
	// RETURN is the return keyword.
	RETURN = "RETURN"
)

//nolint:gochecknoglobals // Immutable keyword lookup table; never mutated after init.
var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns the token type for an identifier.
//
// If the identifier corresponds to a keyword, that token type will be returned (e.g. "fn" maps to FUNCTION); any other
// identity yields the IDENT type, meaning the lexer treats it as a user-defined variable of function name. The lookup
// is case-sensitive.
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

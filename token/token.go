package token

// Using string is more debug-friendly because we can print strings
// Using strings might not be as performant as Bytes and Ints etc.
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special characters
	ILLEGAL = "ILLEGAL" // Unknown character/token
	EOF     = "EOF"     // End of file, marks the end of the parser

	// Identifiers and literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1, 2, 3, 10, 50, ...

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

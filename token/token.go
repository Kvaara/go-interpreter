/*
Lexer is going to output these tokens as a part of the Lexical Analysis.
This process is known as Lexing and is done by a Lexer (also known as Tokenizer/Scanner).

Lexer is going to be given an input. For example: "let x = 5 + 5;"
This input will be processed and categorized into the following data structure that holds a list of Tokens:
	[
		LET,
		IDENTIFIER("x"),
		EQUAL_SIGN,
		INTEGER(5),
		PLUS_SIGN,
		INTEGER(5),
		SEMICOLON
	]

These tokens will then be fed into the Parser, which turns the tokens into an Abstract Syntax Tree (AST).

Source code => Tokens => AST
*/

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

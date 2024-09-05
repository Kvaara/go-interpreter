# go-interpreter

An interpreter written in Go. Courtesy of Thorsten Ball's book "Writing An Interpreter In Go"

The basic What-and-How process of an Interpreter is as follows:

- [Lexical Analysis](#lexical-analysis) => [Parser](#parser) => [Evaluation](#evaluation)

  - Source code => Tokens => Parser => AST => Evaluation

## The Components of An Interpreter

### Lexical Analysis

In programming, Lexical Analysis is **the process where source code is read, processed, and semantically categorized into a data structure, which holds Tokens.**

This process is known as Lexing and is **done by a Lexer** (*also known as a Tokenizer/Scanner*).

For example, if the following input `let x = 5 + 5;` were to be fed into a Lexer, it would be processed into a list of Tokens:

```go
[
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]
```

**Tokens are the smallest unit of meaning in source code.**

As part of the Lexical Analysis, these **tokens are then going to be fed into the [Parser](#parser)** which turns the tokens into an [Abstract Syntax Tree (AST)](#abstract-tree-syntax-ast).

### Parser

#### Abstract Tree Syntax (AST)

### Evaluation

## Improvements

### Debug-Friendliness

In a production environment it would make sense to attach filenames and line numbers to tokens, to better track down lexing and parsing errors. So it would be better to initialize the lexer with an `io.Reader` and the filename.

### Full Unicode Support

The [Lexer](#lexical-analysis) supports only ASCII characters for simplicity. However, the interpreter could be made to support the full Unicode and UTF-8 range. Implementing full Unicode support would also enable the use of emojis!

### Allow More Characters In Identifiers

We could continue the `isLetter` function in [lexer.go](/lexer/lexer.go) to allow more characters in identifiers (e.g., `this!is!a!function` or `exists?`).

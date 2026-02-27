package token

type TokenKind string

// keywords for the lox language
var Keywords = []string{"and", "class", "else", "false", "for", "fun", "if", "nil", "or", "print", "return", "super", "this", "true", "var", "while"}

const (
	// KEYWORDS
	AND    TokenKind = "AND"
	CLASS  TokenKind = "CLASS"
	ELSE   TokenKind = "ELSE"
	FALSE  TokenKind = "FALSE"
	FOR    TokenKind = "FOR"
	FUN    TokenKind = "FUN"
	IF     TokenKind = "IF"
	NIL    TokenKind = "NIL"
	OR     TokenKind = "OR"
	PRINT  TokenKind = "PRINT"
	RETURN TokenKind = "RETURN"
	SUPER  TokenKind = "SUPER"
	THIS   TokenKind = "THIS"
	TRUE   TokenKind = "TRUE"
	VAR    TokenKind = "VAR"
	WHILE  TokenKind = "WHILE"

	// Literals, String, Number, Symbols, Espace Seq
	IDENTIFIER      TokenKind = "IDENTIFIER"
	LEFT_PAREN      TokenKind = "LEFT_PAREN"
	RIGHT_PAREN     TokenKind = "RIGHT_PAREN"
	LEFT_BRACE      TokenKind = "LEFT_BRACE"
	RIGHT_BRACE     TokenKind = "RIGHT_BRACE"
	STAR            TokenKind = "STAR"
	DOT             TokenKind = "DOT"
	COMMA           TokenKind = "COMMA"
	SEMI_COLON      TokenKind = "SEMI_COLON"
	PLUS            TokenKind = "PLUS"
	MINUS           TokenKind = "MINUS"
	EQUAL           TokenKind = "EQUAL"
	EQUAL_EQUAL     TokenKind = "EQUAL_EQUAL"
	BANG            TokenKind = "BANG"
	BANG_EQUAL      TokenKind = "BANG_EQUAL"
	LESS_EQUAL      TokenKind = "LESS_EQUAL"
	GREATER_EQUAL   TokenKind = "GREATER_EQUAL"
	LESS            TokenKind = "LESS"
	GREATER         TokenKind = "GREATER"
	SLASH           TokenKind = "SLASH"
	STRING          TokenKind = "STRING"
	NUMBER          TokenKind = "NUMBER"
	NEW_LINE        TokenKind = "NEW_LINE"
	TAB             TokenKind = "TAB"
	CARRIAGE_RETURN TokenKind = "CARRIAGE_RETURN"
	SPACE           TokenKind = "SPACE"
)

type Token struct {
	Kind    TokenKind
	Lex     []byte
	Literal string
}

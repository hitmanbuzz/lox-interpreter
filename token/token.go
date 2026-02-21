package token

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	STAR
	DOT
	COMMA
	SEMI_COLON
	PLUS
	MINUS
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS_EQUAL
	GREATER_EQUAL
	LESS
	GREATER
	SLASH
	NULL
)

type Token struct {
	Kind      TokenType
	KindValue string
	Lex       string
	Literal   string
}

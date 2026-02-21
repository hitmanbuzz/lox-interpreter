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
	SLASH
	NULL
)

type TokenData struct {
	Kind      TokenType
	KindValue string
	Lex       string
	Literal   string
}

type Token struct {
	Prev TokenData
	Curr TokenData
}

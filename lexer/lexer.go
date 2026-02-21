package lexer

import (
	"fmt"
	"mylang/token"
)

type Lexer struct {
	tokens []token.Token
	// store the current lexing line number
	lineNo   uint
	ExitCode int
}

func NewLexer() *Lexer {
	return &Lexer{
		tokens:   make([]token.Token, 0),
		lineNo:   1,
		ExitCode: 0,
	}
}

func (l *Lexer) Tokenize(char string, currIdx uint, str string) uint {
	nt := token.Token{
		Prev: token.TokenData{
			Kind:      token.NULL,
			KindValue: "",
			Lex:       "",
			Literal:   "null",
		},
		Curr: token.TokenData{
			Kind:      token.NULL,
			KindValue: "",
			Lex:       "",
			Literal:   "null",
		},
	}

	if currIdx >= 1 {
		prevChar := string(str[currIdx-1])
		pLex, pKind, pKindValue, _ := l.MatchToken(prevChar, "")
		nt.Prev.Lex = pLex
		nt.Prev.Kind = pKind
		nt.Prev.KindValue = pKindValue
	}

	nextChar := string(str[currIdx+1])
	lex, kind, kindValue, jump := l.MatchToken(char, nextChar)
	if kind != token.NULL {
		nt.Curr.Lex = lex
		nt.Curr.Kind = kind
		nt.Curr.KindValue = kindValue
		l.tokens = append(l.tokens, nt)
	}

	return jump
}

func (l *Lexer) MatchToken(c string, nc string) (string, token.TokenType, string, uint) {
	switch c {
	// skip if it is whitespace
	case " ":
		return c, token.NULL, "", 1
	// count the line number with \n
	case "\n":
		l.lineNo += 1
		return c, token.NULL, "", 1
	case "(":
		return c, token.LEFT_PAREN, "LEFT_PAREM", 1
	case ")":
		return c, token.RIGHT_PAREN, "RIGHT_PAREN", 1
	case "{":
		return c, token.LEFT_BRACE, "LEFT_BRACE", 1
	case "}":
		return c, token.RIGHT_BRACE, "RIGHT_BRACE", 1
	case ",":
		return c, token.COMMA, "COMMA", 1
	case ".":
		return c, token.DOT, "DOT", 1
	case "=":
		if nc == "=" {
			return "==", token.EQUAL_EQUAL, "EQUAL_EQUAL", 2
		}
		return c, token.EQUAL, "EQUAL", 1
	case "+":
		return c, token.PLUS, "PLUS", 1
	case "-":
		return c, token.MINUS, "MINUS", 1
	case ";":
		return c, token.SEMI_COLON, "SEMI_COLON", 1
	case "/":
		return c, token.SLASH, "SLASH", 1
	case "*":
		return c, token.STAR, "STAR", 1
	default:
		fmt.Printf("[line %d] Error: Unexpected character: %s\n", l.lineNo, c)
		l.ExitCode = 65
		return c, token.NULL, "", 1
	}
}

func (l *Lexer) Display() {
	for _, t := range l.tokens {
		fmt.Printf("%s %s %s\n", t.Curr.KindValue, t.Curr.Lex, t.Curr.Literal)
	}

	fmt.Println("EOF  null")
}

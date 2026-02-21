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
		Kind:      token.NULL,
		KindValue: "",
		Lex:       "",
		Literal:   "null",
	}

	nextChar := string(str[currIdx+1])
	lex, kind, kindValue, jump := l.MatchToken(char, nextChar)
	if kind != token.NULL {
		nt.Lex = lex
		nt.Kind = kind
		nt.KindValue = kindValue
		l.tokens = append(l.tokens, nt)
	}

	return jump
}

// @param:
//
// c -> current char
//
// nc -> next char
//
// @return
//
// 1st (string) -> lex
//
// 2nd (TokenType) -> TokenType
//
// 3rd (uint) -> Jump Amount
//
// The 3rd return value is the number of indices to skip from the current char.
// Default is always 1
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
		return c, token.LEFT_PAREN, "LEFT_PAREN", 1
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
	case "!":
		if nc == "=" {
			return "!=", token.BANG_EQUAL, "BANG_EQUAL", 2
		}
		return "", token.NULL, "", 1
	case ">":
		if nc == "=" {
			return ">=", token.GREATER_EQUAL, "GREATER_EQUAL", 2
		}
		return ">", token.GREATER, "GREATER", 1
	case "<":
		if nc == "=" {
			return "<=", token.LESS_EQUAL, "LESS_EQUAL", 2
		}
		return "<", token.LESS, "LESS", 1
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
		fmt.Printf("%s %s %s\n", t.KindValue, t.Lex, t.Literal)
	}

	fmt.Println("EOF  null")
}

package lexer

import (
	"fmt"
	"mylang/token"
	"mylang/utils"
)

type Lexer struct {
	tokens   []token.Token
	CurrIdx  int
	LineNo   int
	Source   string
	ExitCode int
}

func NewLexer() *Lexer {
	return &Lexer{
		tokens:   make([]token.Token, 0),
		CurrIdx:  0,
		LineNo:   1,
		ExitCode: 0,
	}
}

func (l *Lexer) Tokenize() int {
	nextByte := l.Source[l.CurrIdx+1]
	l.MatchToken(nextByte)
	return l.CurrIdx
}

func (l *Lexer) MatchToken(nextByte byte) {
	currByte := l.peek()

	switch currByte {
	case ' ':
		l.addToken(token.SPACE, utils.ToByteArr(currByte), "<SPACE>")
		l.advance(1)
	case '\t':
		l.addToken(token.SPACE, utils.ToByteArr(currByte), "\\t")
		l.advance(1)
	case '\r':
		l.addToken(token.CARRIAGE_RETURN, utils.ToByteArr(currByte), "\\r")
		l.advance(1)
	case '\n':
		l.addToken(token.NEW_LINE, utils.ToByteArr(currByte), "null")
		l.advance(1)
		l.incrementLine()
	case '(':
		l.addToken(token.LEFT_PAREN, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case ')':
		l.addToken(token.RIGHT_PAREN, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '{':
		l.addToken(token.LEFT_BRACE, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '}':
		l.addToken(token.RIGHT_BRACE, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case ',':
		l.addToken(token.COMMA, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '.':
		l.addToken(token.DOT, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '=':
		if nextByte == '=' {
			l.advance(2)
			l.addToken(token.EQUAL_EQUAL, utils.ToByteArr(currByte), "null")
		} else {
			l.addToken(token.EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '!':
		if nextByte == '=' {
			l.addToken(token.BANG_EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(2)
		} else {
			l.addToken(token.BANG, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '>':
		if nextByte == '=' {
			l.addToken(token.GREATER_EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(2)
		} else {
			l.addToken(token.GREATER, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '<':
		if nextByte == '=' {
			l.addToken(token.LESS_EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(2)
		} else {
			l.addToken(token.LESS, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '+':
		l.addToken(token.PLUS, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '-':
		l.addToken(token.MINUS, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case ';':
		l.addToken(token.SEMI_COLON, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '/':
		if nextByte == '/' {
			l.scanComment()
		} else {
			l.addToken(token.SLASH, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '"':
		str, isString := l.scanString()
		if isString {
			l.addToken(token.STRING, str, string(str))
		} else {
			fmt.Printf("[line %d] Error: Unterminated string.\n", l.LineNo)
		}
	case '*':
		l.addToken(token.STAR, utils.ToByteArr(currByte), "null")
		l.advance(1)
	default:
		fmt.Printf("[line %d] Error: Unexpected character: %c\n", l.LineNo, currByte)
		l.ExitCode = 65
		l.advance(1)
	}
}

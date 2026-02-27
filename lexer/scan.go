package lexer

import (
	"mylang/token"
	"mylang/utils"
	"slices"
)

func (l *Lexer) scanComment() {
	for {
		if l.isAtEnd() {
			break
		} else if l.peek() == '\n' {
			l.Line++
			l.advance(1)
			break
		}
		l.advance(1)
	}
}

func (l *Lexer) scanString() ([]byte, bool) {
	isString := false
	var str []byte
	l.advance(1)

	for {
		if l.isAtEnd() {
			break
		} else if l.peek() == '\n' {
			l.Line++
		} else if l.peek() == '"' {
			isString = true
			l.advance(1)
			break
		}

		str = append(str, l.peek())
		l.advance(1)
	}

	return str, isString
}

func (l *Lexer) scanNumber() []byte {
	var nums []byte
	isDot := false
	isNumberStart := false

	for {
		if l.isAtEnd() {
			break
		} else if l.peek() == '\n' {
			l.Line++
			l.advance(1)
			break
		}

		nextByte := l.Source[l.CurrIdx+1]

		if utils.IsNum(l.peek()) {
			nums = append(nums, l.peek())
			isNumberStart = true
		} else if l.peek() == '.' && utils.IsNum(nextByte) && !isDot && isNumberStart {
			isDot = true
			nums = append(nums, l.peek())
		} else {
			isNumberStart = false
			break
		}

		l.advance(1)
	}

	lastB := nums[len(nums)-1]
	if lastB != '.' && !isDot {
		nums = append(nums, '.')
		nums = append(nums, '0')
	}

	return nums
}

func (l *Lexer) scanIdentifier() ([]byte, bool) {
	var iden []byte

	for {
		if l.isAtEnd() {
			break
		} else if l.peek() == '\n' {
			l.advance(1)
			l.Line++
			break
		}

		if l.peek() == ' ' {
			l.advance(1)
			break
		}

		iden = append(iden, l.peek())
		l.advance(1)
	}

	strValue := string(iden)
	if slices.Contains(token.Keywords, strValue) {
		return iden, false
	}

	return iden, true
}

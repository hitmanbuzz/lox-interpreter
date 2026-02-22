package lexer

func (l *Lexer) scanComment() {
	for {
		s := l.peek()
		if l.isAtEnd() {
			break
		}

		if s == '\n' {
			l.incrementLine()
			l.advance(1)
			break
		}
		l.advance(1)
	}
}

func (l *Lexer) scanString() ([]byte, bool) {
	isString := false
	var str []byte
	skip := 0

	for {
		if skip == 0 {
			l.advance(1)
			skip++
			continue
		}

		if l.isAtEnd() {
			break
		}

		if l.peek() == '"' {
			isString = true
			l.advance(1)
			skip++
			break
		}

		if l.peek() == '\n' {
			l.incrementLine()
		}

		str = append(str, l.peek())
		l.advance(1)
		skip++
	}

	return str, isString
}

package main

import (
	"fmt"
	"mylang/lexer"
	"mylang/utils"
	"os"
)

func main() {
	lexer := lexer.NewLexer()
	fname := "test.lox"

	if !utils.IsFileExist(fname) {
		fmt.Printf("File `%s` doesn't exist\n", fname)
		return
	}

	data, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println("EOF  null")
		return
	}

	d := string(data)
	var idx uint = 0

	for {
		if idx+1 <= uint(len(d))-1 {
			jump := lexer.Tokenize(string(d[idx]), idx, d)
			idx += jump
		} else {
			break
		}
	}

	lexer.Display()
	os.Exit(lexer.ExitCode)
}

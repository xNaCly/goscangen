package main

import "strings"

var test = `
; comment
[Section]
1key123=value_12
`

func main() {
	lexer := Lexer{Builder: &strings.Builder{}}
	err := lexer.NewInput(test)
	if err != nil {
		panic(err)
	}
	tok, err := lexer.Lex()
	if err != nil {
		panic(err)
	}
	Debug(tok)
}

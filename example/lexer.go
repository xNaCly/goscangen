package main

import (
	"errors"
	"fmt"
	"strings"
)

type TokenType uint

type Token struct {
	Pos  int
	Line int
	Type TokenType
	Raw  string
}

var KEYWORDS = map[string]TokenType{}

const (
	UNKNOWN TokenType = iota
	EOF

	// generated
	braket_left
	braket_right
	equal
	ident
)

var LOOKUP = map[TokenType]string{
	UNKNOWN: "UNKNOWN",
	EOF:     "EOF",
	// generated
	braket_left:  "braket_left",
	braket_right: "braket_right",
	equal:        "equal",
	ident:        "ident",
}

func Debug(token []Token) {
	fmt.Printf("| %4s | %3s | %20s | %50s |\n", "line", "pos", "type", "raw")
	fmt.Printf("| %4s | %3s | %20s | %50s |\n", "-", "-", "-", "-")
	for _, t := range token {
		fmt.Printf("| %04d | %03d | %20s | %50s |\n", t.Line, t.Pos, LOOKUP[t.Type], t.Raw)
	}
}

type Lexer struct {
	pos     int
	lPos    int
	l       int
	in      string
	inL     int
	cc      byte
	Builder *strings.Builder
}

func (l *Lexer) NewInput(input string) error {
	l.Builder.Reset()
	l.pos = 0
	l.lPos = 0
	l.in = input
	l.l = 0
	l.inL = len(input)
	if l.inL == 0 {
		return errors.New("can't accept empty input")
	}
	l.cc = input[0]
	return nil
}

func (l *Lexer) advance() {
	if l.pos+1 < l.inL {
		l.pos++
		l.lPos++
		l.cc = l.in[l.pos]
	} else {
		l.cc = 0
	}
}

// generated [\a \d _]*
func charMatchesIdent(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '_'
}

// generated [\a \d _]*
func (l *Lexer) ident() (TokenType, string) {
	for charMatchesIdent(l.cc) {
		l.Builder.WriteByte(l.cc)
		l.advance()
	}
	str := l.Builder.String()
	t := ident
	if tt, ok := KEYWORDS[str]; ok {
		t = tt
	}
	l.Builder.Reset()
	return t, str
}

func (l *Lexer) Lex() ([]Token, error) {
	r := make([]Token, 0)
	for l.cc != 0 {
		t := UNKNOWN
		switch l.cc {
		case '\n': // skip newlines and increment line counter
			l.l++
			l.lPos = 0
			l.advance()
			continue
		case ' ', '\t': // skip whitespace
			l.advance()
			continue
		case ';': // generated
			for l.cc != 0 && l.cc != '\n' {
				l.advance()
			}
			continue
		case '[': // generated
			t = braket_left
		case ']': // generated
			t = braket_right
		case '=': // generated
			t = equal
		default:
			// generated
			if charMatchesIdent(l.cc) {
				tt, str := l.ident()
				r = append(r, Token{
					Pos:  l.lPos - len(str),
					Line: l.l,
					Type: tt,
					Raw:  str,
				})
				continue
			}
		}
		r = append(r, Token{
			Pos:  l.lPos,
			Line: l.l,
			Type: t,
			Raw:  "",
		})
		l.advance()
	}
	return r, nil
}

package core

import (
	_ "embed"
	"log"
	"os"
	"text/template"
)

//go:embed template/LexerTemplate.txt
var LexerTemplate string

type SourceData struct {
	SourcePath string
	DestPath   string
	Package    string
}

type CompileData struct {
	SourceData
	TokenTypes []string
	KeyWords   []string
}

func Start(c SourceData) {
	tmpl, err := template.New("lexerTemplate").Parse(LexerTemplate)
	if err != nil {
		log.Fatalln(err)
	}
	// TEST: remove once lexer and parser are done.
	data := CompileData{
		SourceData: c,
		TokenTypes: []string{"braket_left", "braket_right", "equal", "ident"},
		KeyWords:   []string{},
	}
	tmpl.Execute(os.Stdout, data)
}

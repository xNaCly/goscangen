package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/xnacly/goscangen/core"
)

func main() {
	i := flag.String("i", "", "location of scangen definition")
	o := flag.String("o", fmt.Sprintf("Lexer_%s.go", time.Now().Format("2006-01-02_15-04-05")), "file to write generated code to")
	d := flag.Bool("debug", false, "enable file and line numbers in logs")
	p := flag.String("p", "main", "define package of lexer ")

	flag.Parse()
	if *d {
		log.SetFlags(log.Lshortfile)
	} else {
		log.SetPrefix("scangen: ")
		log.SetFlags(0)
	}
	if *i == "" {
		log.Fatalln("missing or empty scangen definition argument")
	}
	core.Start(core.SourceData{
		SourcePath: *i,
		DestPath:   *o,
		Package:    *p,
	})
}

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
	o := flag.String("o", "", "destination to write Lexer to")
	d := flag.Bool("debug", false, "enable file and line numbers in logs")
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
	if *o == "" {
		t := fmt.Sprintf("Lexer_%s.go", time.Now().Format("2006-01-02_15-04-05"))
		log.Printf("got no value for '-o', falling back to %q", t)
		*o = t
	}
	core.Start(*i, *o)
}

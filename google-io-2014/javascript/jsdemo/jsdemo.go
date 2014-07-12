package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sourcegraph/talks/google-io-2014/javascript"
)

func main() {
	flag.Parse()
	var file string
	if flag.NArg() == 1 {
		file = flag.Arg(0)
	} else {
		file = "google-io-2014/javascript/sample.js"
	}
	log.SetFlags(0)
	log.Println("Analyzing", file, "for funcs & vars")
	// START OMIT
	defs, _, err := (&javascript.JSAnalyzer{}).Analyze(file)
	// END OMIT
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NAME      \tTYPE")
	fmt.Println("----      \t----")
	for _, def := range defs {
		fmt.Printf("%-10s\t%s\n", def.Name, def.Type)
	}
}

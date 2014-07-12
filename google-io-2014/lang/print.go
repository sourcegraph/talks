package lang

import "fmt"

func PrintHandlers() {
	fmt.Println("LANG      \tANALYZER")
	fmt.Println("----      \t--------")
	for lang, a := range analyzers {
		fmt.Printf("%-10s\t%T\n", lang, a)
	}
}

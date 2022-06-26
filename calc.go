package main

import (
	"fmt"
	"github.com/jmptc/carner-calc/lexer"
	"github.com/jmptc/carner-calc/parser"
	"github.com/jmptc/carner-calc/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to carner-calc")
	input := ""
	//if len(os.Args) == 2 {
	if input != "" {
		l := lexer.NewLexer(input)
		//fmt.Println(l.GetTokens())
		p := parser.New(l.GetTokens())
		fmt.Println(p.Parse())
	} else {
		repl.Repl(os.Stdin, os.Stdout)
	}
}

package main

import (
    "fmt"
    "github.com/jmptc/carner-calc/repl"
    "os"
)

func main() {
    fmt.Println("Welcome to carner-calc")
    /*
    input := "(1+2)*(3-4)/10"
    fmt.Println("input: ", input)
    l := lexer.NewLexer(input)
    fmt.Println(l.GetTokens())
    */
    repl.Repl(os.Stdin, os.Stdout) 
}

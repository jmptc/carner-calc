package main

import (
    "fmt"
    "github.com/jmptc/carner-calc/lexer"
)

func main() {
    fmt.Println("Welcome to carner-calc")
    input := "(1+2)*(3-4)/10"
    fmt.Println("input: ", input)
    l := lexer.NewLexer(input)
    fmt.Println(l.GetTokens())
}

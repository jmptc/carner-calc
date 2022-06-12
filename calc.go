package main

import (
    "fmt"
    "github.com/jmptc/carner-calc/lexer"
)

func main() {
    fmt.Println("Welcome to carner-calc")

    l := lexer.NewLexer("()")
    fmt.Println(l.GetTokens())
}

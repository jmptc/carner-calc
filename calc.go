package main

import (
    "fmt"
    "github.com/jmptc/carner-calc/lexer"
    "github.com/jmptc/carner-calc/repl"
    "github.com/jmptc/carner-calc/parser"
    "os"
)

func main() {
    fmt.Println("Welcome to carner-calc")
    input := "1 + 2 * 3"
    //if len(os.Args) == 2 {
    if input != "" {
         l := lexer.NewLexer(input)
        //fmt.Println(l.GetTokens())
        p := parser.New(l.GetTokens())

        fmt.Printf("%s\n", p.CurTok())
        for p.Advance() {
            fmt.Printf("%s\n", p.CurTok())
        }
    } else {
        repl.Repl(os.Stdin, os.Stdout) 
    }
}

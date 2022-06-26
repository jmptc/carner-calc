package repl

import (
    "fmt"
    "io"
    "bufio"

    "github.com/jmptc/carner-calc/lexer"
    "github.com/jmptc/carner-calc/parser"
)

func Repl(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for scanner.Scan() {
        l := lexer.NewLexer(scanner.Text())
        tokens := l.GetTokens()
        
        p := parser.New(tokens)
        fmt.Fprintf(out, "%s\n", p.Parse())
    }
}

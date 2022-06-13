package repl

import (
    "io"
    "bufio"
    "github.com/jmptc/carner-calc/lexer"
    "fmt"
)

func Repl(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for scanner.Scan() {
        l := lexer.NewLexer(scanner.Text())
        tokens := l.GetTokens()

        fmt.Fprintf(out, "%s\n", tokens)
    }
}

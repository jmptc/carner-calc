package lexer

import (
    "testing"
    "github.com/jmptc/carner-calc/token"
)

func TestGetTokens(t *testing.T) {
    input := `()+*-/`

    tests := []struct {
        expectedType    string
        expectedValue   string
    }{
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.PLUS, "+"},
        {token.ASTERISK, "*"},
        {token.MINUS, "-"},
        {token.SLASH, "/"},
        {token.EOF, token.EOF},
    }

    l := Lexer{input: input}
    tokens := l.GetTokens()

    if len(tokens) != 7 {
        t.Fatalf("Mismatched length between tests and lexer tokens got: %d expected: %d", len(tokens), 7)
    }

    for i, testTok := range tests {
        if tokens[i].Type != testTok.expectedType {
            t.Fatalf("tests[%d] failed. Wrong token type got: %s expected: %s", i, tokens[i].Type, testTok.expectedType)
        }
    }
}

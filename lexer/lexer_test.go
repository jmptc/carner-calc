package lexer

import (
    "testing"
    "github.com/jmptc/carner-calc/token"
)

func TestGetTokens(t *testing.T) {
    input := `(254*13) / ( 14 + 1 - 3)`

    tests := []struct {
        expectedType    string
        expectedValue   string
    }{
        {token.LPAREN, "("},
        {token.NUM, "254"},
        {token.ASTERISK, "*"},
        {token.NUM, "13"},
        {token.RPAREN, ")"},
        {token.SLASH, "/"},
        {token.LPAREN, "("},
        {token.NUM, "14"},
        {token.PLUS, "+"},
        {token.NUM, "1"},
        {token.MINUS, "-"},
        {token.NUM, "3"},
        {token.RPAREN, ")"},
        {token.EOF, token.EOF},
    }

    l := Lexer{input: input}
    tokens := l.GetTokens()

    if len(tokens) != len(tests) {
        t.Fatalf("Mismatched length between tests and lexer tokens got: %d expected: %d", len(tokens), 7)
    }

    for i, testTok := range tests {
        if tokens[i].Type != testTok.expectedType {
            t.Fatalf("tests[%d] failed. Wrong token type got: %s expected: %s", i, tokens[i].Type, testTok.expectedType)
        }
    }
}

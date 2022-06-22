package parser

import (
    "testing"

    "github.com/jmptc/carner-calc/token"
    "github.com/jmptc/carner-calc/lexer"
)

func TestParserAdvance(t *testing.T) {
    input := "1 + 2"

    expectedTokens := []token.Token{
        {Type: token.NUM, Value: "1"},
        {Type: token.PLUS, Value: "+"},
        {Type: token.NUM, Value: "2"},
        {Type: token.EOF, Value: token.EOF},
    }


    l := lexer.NewLexer(input)
    p := New(l.GetTokens())

    for _, expectedToken := range expectedTokens {
        if !compareTokens(expectedToken, p.curTok) {
            t.Errorf("Token mismatch. Expected %s, Got %s\n", expectedToken, p.curTok)
        }

        p.advance()
    }
}

func compareTokens(a, b token.Token) bool {
    return a.Type == b.Type && a.Value == b.Value
}


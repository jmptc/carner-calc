package parser

import (
	"testing"

	"github.com/jmptc/carner-calc/lexer"
	"github.com/jmptc/carner-calc/token"
)

func TestExpressionParsing(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			"1 + 2",
			"(1 + 2)",
		},
		{
			"3 * 4 - 1",
			"((3 * 4) - 1)",
		},
		{
			"2 / 3 * 4",
			"((2 / 3) * 4)",
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l.GetTokens())

		output := p.Parse()

		if output != tt.output {
			t.Errorf("Output mismatch. Expected %s, got %s\n", tt.output, output)
		}
	}
}

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

		p.Advance()
	}
}

func compareTokens(a, b token.Token) bool {
	return a.Type == b.Type && a.Value == b.Value
}

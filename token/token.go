package token

import "fmt"

const (
	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
	ASTERISK = "*"
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	NUM      = "NUM"
	EOF      = "EOF"
)

type Token struct {
	Type  string
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Value: %s}", t.Type, t.Value)
}

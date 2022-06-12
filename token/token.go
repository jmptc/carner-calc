package token

const (
    LPAREN      = "LPAREN"
    RPAREN      = "RPAREN"
    ASTERISK    = "*"
    PLUS        = "+"
    MINUS       = "-"
    SLASH       = "/"
    NUM         = "NUM"
    EOF         = "EOF"
)

type Token struct {
    Type    string
    Value   string
}

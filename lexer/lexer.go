package lexer

import (
    "github.com/jmptc/carner-calc/token"
)

type Lexer struct {
    input   string
    pos     int
}

func NewLexer(input string) Lexer {
    return Lexer{input: input, pos: 0}
}

func (l *Lexer) GetTokens() []token.Token {
    tokens := []token.Token{}

    for {
        tok := l.getToken()
        tokens = append(tokens, tok)

        if tok.Type == token.EOF {
            return tokens
        }
        
    }
}

func (l *Lexer) getToken() token.Token {
    if l.pos >= len(l.input) {
        return token.Token{Type: token.EOF, Value: token.EOF}
    }
    ch := l.input[l.pos]

    switch ch {
    case '(':
        return l.newToken(ch, token.LPAREN)
    case ')':
        return l.newToken(ch, token.RPAREN)
    case '*':
        return l.newToken(ch, token.ASTERISK)
    case '-':
        return l.newToken(ch, token.MINUS)
    case '/':
        return l.newToken(ch, token.SLASH)
    case '+':
        return l.newToken(ch, token.PLUS)
    default:
        return token.Token{Type: token.EOF, Value: token.EOF}
    }
}

func (l *Lexer) newToken(ch byte, tokenType string) token.Token {
    l.pos += 1
    return token.Token{Type: tokenType, Value: string(ch)}
}



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
    l.skipWhitespace()

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
        if isDigit(ch) {
            return l.number()
        }
        return token.Token{Type: token.EOF, Value: token.EOF}
    }
}

func (l *Lexer) skipWhitespace() {
    for !l.isNextEOF() {
        if !isWhitespace(l.input[l.pos]) {
            break
        } else {
            l.pos++
        }
    }
}

func (l *Lexer) number() token.Token {
    start := l.pos

    for !l.isNextEOF() {
        if isDigit(l.peek()) {
            l.pos++
        } else {
            break;
        }
    }

    tok := token.Token{Type: token.NUM, Value: l.input[start:l.pos+1]}
    l.pos++
    return tok
}

func (l *Lexer) peek() byte {
        return l.input[l.pos + 1] 
}

func (l *Lexer) isNextEOF() bool {
    if l.pos  >= len(l.input) - 1 {
        return true
    }
    return false
}

func (l *Lexer) newToken(ch byte, tokenType string) token.Token {
    l.pos += 1
    return token.Token{Type: tokenType, Value: string(ch)}
}

func isDigit(ch byte) bool {
    return ch >= 48 && ch <= 57
}

func isLetter(ch byte) bool {
    return (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122)
}

func isWhitespace(ch byte) bool {
    return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' 
}

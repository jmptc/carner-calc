package parser

import (
    //"fmt"
    //"strconv"

    "github.com/jmptc/carner-calc/token"
    "github.com/jmptc/carner-calc/ast"
)

const (
    LOWEST = iota + 1
    SUM
    PRODUCT
    GROUP
)

type prefixParseFunc func() ast.Expression
type infixParseFunc func(ast.Expression) ast.Expression

type Parser struct {
    tokens      []token.Token
    curTokIdx   int
    curTok      token.Token
    peekTok     token.Token

    prefixParseFuncs map[string]prefixParseFunc
    infixParseFuncs  map[string]infixParseFunc
}

func New(tokens []token.Token) *Parser {
    p := &Parser{tokens: tokens}

    // init tokens
    p.Advance()
    p.Advance()

    return p
}

func (p *Parser) Advance() bool {
    if p.curTok.Type != token.EOF {
        p.curTok = p.peekTok
        if p.curTokIdx < len(p.tokens) {
            p.peekTok = p.tokens[p.curTokIdx]
            p.curTokIdx += 1
        }
        return true
    } else {
        return false
    }
}

func (p *Parser) CurTok() token.Token {
    return p.curTok
}
    


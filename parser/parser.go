package parser

import (
	"fmt"
	"strconv"

	"github.com/jmptc/carner-calc/ast"
	"github.com/jmptc/carner-calc/token"
)

const (
	LOWEST = iota + 1
	SUM
	PRODUCT
	GROUP
)

var precedenceMap = map[string]int{
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.ASTERISK: PRODUCT,
	token.SLASH:    PRODUCT,
    token.LPAREN:   GROUP,
}

type prefixParseFunc func() ast.Expression
type infixParseFunc func(ast.Expression) ast.Expression

type Parser struct {
	tokens    []token.Token
	curTokIdx int
	curTok    token.Token
	peekTok   token.Token

	errors           []string
	prefixParseFuncs map[string]prefixParseFunc
	infixParseFuncs  map[string]infixParseFunc
}

func New(tokens []token.Token) *Parser {
	p := &Parser{tokens: tokens}
	p.errors = []string{}

	p.prefixParseFuncs = make(map[string]prefixParseFunc)
	p.prefixParseFuncs[token.NUM] = p.parseNumLiteral
    p.prefixParseFuncs[token.LPAREN] = p.parseGroup

	p.infixParseFuncs = make(map[string]infixParseFunc)
	p.infixParseFuncs[token.PLUS] = p.parseInfixExpression
	p.infixParseFuncs[token.MINUS] = p.parseInfixExpression
	p.infixParseFuncs[token.ASTERISK] = p.parseInfixExpression
	p.infixParseFuncs[token.SLASH] = p.parseInfixExpression

	// init tokens
	p.Advance()
	p.Advance()

	return p
}

func (p *Parser) Parse() string {
	expression := p.parseExpression(LOWEST)
	return expression.String()
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	// 1 + 2 * 3
	prefix := p.prefixParseFuncs[p.curTok.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.EOF) && precedence < p.peekPrecedence() {
		infix := p.infixParseFuncs[p.peekTok.Type]

		p.Advance()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseInfixExpression(leftExp ast.Expression) ast.Expression {
	binaryExp := &ast.BinaryExpression{
		LeftExp: leftExp,
		Op:      p.curTok,
	}

	precedence := p.curPrecedence()
	p.Advance()
	binaryExp.RightExp = p.parseExpression(precedence)

	return binaryExp
}

func (p *Parser) parseNumLiteral() ast.Expression {
	num := &ast.NumLiteral{Token: p.curTok}

	val, err := strconv.Atoi(p.curTok.Value)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Error parsing %s to integer in parseNumLiteral. Token: %s", p.curTok.Value, p.curTok))
		return nil
	}

	num.Value = val

	return num
}

func (p *Parser) parseGroup() ast.Expression {
    p.Advance()

    exp := p.parseExpression(LOWEST)

    if !p.peekTokenIs(token.RPAREN) {
        return nil
    }

    p.Advance()

    return exp
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

func (p *Parser) curPrecedence() int {
	if prec, ok := precedenceMap[p.curTok.Type]; ok {
		return prec
	}

	return LOWEST
}

func (p *Parser) CurTok() token.Token {
	return p.curTok
}

func (p *Parser) peekTokenIs(tokenType string) bool {
	return p.peekTok.Type == tokenType
}

func (p *Parser) peekPrecedence() int {
	return precedenceMap[p.peekTok.Type]
}

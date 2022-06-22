package ast

import (
    "github.com/jmptc/carner-calc/token"
)

type Expression interface {
    Type() string
}

type NumLiteral struct {
    Token   token.Token
    Value   int
}

func (ne *NumLiteral) Type() string { return "NumLiteral" }

type BinaryExpression struct {
    leftExp     Expression
    op          token.Token
    rigthExp    Expression
}

func (be *BinaryExpression) Type() string { return "BinaryExpression" }

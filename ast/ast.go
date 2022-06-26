package ast

import (
	"bytes"

	"github.com/jmptc/carner-calc/token"
)

type Expression interface {
	Type() string
	String() string
}

type NumLiteral struct {
	Token token.Token
	Value int
}

func (ne *NumLiteral) Type() string { return "NumLiteral" }

func (ne *NumLiteral) String() string { return ne.Token.Value }

type BinaryExpression struct {
	LeftExp  Expression
	Op       token.Token
	RightExp Expression
}

func (be *BinaryExpression) Type() string { return "BinaryExpression" }

func (be *BinaryExpression) String() string {
	var buf bytes.Buffer

	buf.WriteString("(")
	buf.WriteString(be.LeftExp.String())
	buf.WriteString(" " + be.Op.Value + " ")
	buf.WriteString(be.RightExp.String())
	buf.WriteString(")")

	return buf.String()

}

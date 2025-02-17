package ast

import "github.com/Kresqle/genuml/src/lexer"

// ===================
// LITERAL EXPRESSIONS
// ===================

type NumberExpression struct {
	Value float64
}

func (n NumberExpression) expr() {}

type StringExpression struct {
	Value string
}

func (n StringExpression) expr() {}

type SymbolExpression struct {
	Value string
}

func (n SymbolExpression) expr() {}

// ===================
// COMPLEX EXPRESSIONS
// ===================

type BinaryExpression struct {
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (n BinaryExpression) expr() {}

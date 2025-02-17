package ast

import (
	"github.com/Kresqle/genuml/src/lexer"
)

// --------------------
// Literal Expressions
// --------------------

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

// --------------------
// Complex Expressions
// --------------------

type BinaryExpression struct {
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (n BinaryExpression) expr() {}

type AssignmentExpression struct {
	Assigne       Expression
	AssignedValue Expression
}

func (n AssignmentExpression) expr() {}

type PrefixExpression struct {
	Operator lexer.Token
	Right    Expression
}

func (n PrefixExpression) expr() {}

type MemberExpression struct {
	Member   Expression
	Property string
}

func (n MemberExpression) expr() {}

type CallExpression struct {
	Method    Expression
	Arguments []Expression
}

func (n CallExpression) expr() {}

type ComputedExpression struct {
	Member   Expression
	Property Expression
}

func (n ComputedExpression) expr() {}

type RangeExpression struct {
	Lower Expression
	Upper Expression
}

func (n RangeExpression) expr() {}

type FunctionExpression struct {
	Parameters []Parameter
	Body       []Statement
	ReturnType Type
}

func (n FunctionExpression) expr() {}

type ArrayLiteral struct {
	Contents []Expression
}

func (n ArrayLiteral) expr() {}

type NewExpression struct {
	Instantiation CallExpression
}

func (n NewExpression) expr() {}

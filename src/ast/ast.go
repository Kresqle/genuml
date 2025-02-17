package ast

import "github.com/Kresqle/genuml/src/helpers"

type Statement interface {
	stmt()
}

type Expression interface {
	expr()
}

type Type interface {
	_type()
}

func ExpectExpression[T Expression](expr Expression) T {
	return helpers.ExpectType[T](expr)
}

func ExpectStatement[T Statement](expr Statement) T {
	return helpers.ExpectType[T](expr)
}

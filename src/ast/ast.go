package ast

type Statement interface {
	stmt()
}

type Expression interface {
	expr()
}

type Type interface {
	_type()
}

package ast

type BlockStatement struct {
	Body []Statement
}

func (n BlockStatement) stmt() {}

type ExpressionStatement struct {
	Expression Expression
}

func (n ExpressionStatement) stmt() {}

type VariableDeclarationStatement struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expression
	ExplicitType  Type
}

func (n VariableDeclarationStatement) stmt() {}

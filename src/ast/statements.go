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

type StructProperty struct {
	IsStatic bool // determine whether the property is static
	Type     Type
}

type StructMethod struct {
	IsStatic bool // determine whether the property is static
	Type     Type
}

type StructDeclarationStatement struct {
	StructName string
	Properties map[string]StructProperty
	Methods    map[string]StructMethod
}

func (n StructDeclarationStatement) stmt() {}

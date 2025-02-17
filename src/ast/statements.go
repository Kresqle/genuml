package ast

type BlockStatement struct {
	Body []Statement
}

func (b BlockStatement) stmt() {}

type VarDeclarationStatement struct {
	Identifier    string
	Constant      bool
	AssignedValue Expression
	ExplicitType  Type
}

func (n VarDeclarationStatement) stmt() {}

type ExpressionStatement struct {
	Expression Expression
}

func (n ExpressionStatement) stmt() {}

type Parameter struct {
	Name string
	Type Type
}

type FunctionDeclarationStatement struct {
	Parameters []Parameter
	Name       string
	Body       []Statement
	ReturnType Type
}

func (n FunctionDeclarationStatement) stmt() {}

type IfStatement struct {
	Condition  Expression
	Consequent Statement
	Alternate  Statement
}

func (n IfStatement) stmt() {}

type ImportStatement struct {
	Name string
	From string
}

func (n ImportStatement) stmt() {}

type ForeachStatement struct {
	Value    string
	Index    bool
	Iterable Expression
	Body     []Statement
}

func (n ForeachStatement) stmt() {}

type ClassDeclarationStatement struct {
	Name string
	Body []Statement
}

func (n ClassDeclarationStatement) stmt() {}

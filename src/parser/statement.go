package parser

import (
	"github.com/Kresqle/genuml/src/ast"
	"github.com/Kresqle/genuml/src/lexer"
)

func parse_stmt(p *parser) ast.Statement {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStatement{
		Expression: expression,
	}
}

func parse_var_decl_stmt(p *parser) ast.Statement {
	isConstant := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Inside variable declaration expected to find variable name").Value
	p.expect(lexer.ASSIGNMENT)
	assignedValue := parse_expr(p, assignment)
	p.expect(lexer.SEMI_COLON)

	return ast.VariableDeclarationStatement{
		IsConstant:    isConstant,
		VariableName:  varName,
		AssignedValue: assignedValue,
	}
}

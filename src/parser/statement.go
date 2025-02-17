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
	var explicitType ast.Type
	var assignedValue ast.Expression

	isConstant := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Inside variable declaration expected to find variable name").Value

	// Explicit type could be present
	if p.currentTokenKind() == lexer.COLON {
		p.advance()
		explicitType = parse_type(p, default_bp)
	}

	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		assignedValue = parse_expr(p, assignment)
	} else if explicitType == nil {
		panic("Missing either right-hand side in var declaration of explicit type")
	}

	p.expect(lexer.SEMI_COLON)

	if isConstant && assignedValue == nil {
		panic("Cannot define constant without providing value")
	}

	return ast.VariableDeclarationStatement{
		IsConstant:    isConstant,
		VariableName:  varName,
		AssignedValue: assignedValue,
		ExplicitType:  explicitType,
	}
}

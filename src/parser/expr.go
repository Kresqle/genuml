package parser

import (
	"fmt"
	"strconv"

	"github.com/Kresqle/genuml/src/ast"
	"github.com/Kresqle/genuml/src/lexer"
)

func parse_expr(p *parser, bp binding_power) ast.Expression {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
}

func parse_prefix_expr(p *parser) ast.Expression {
	operatorToken := p.advance()
	expr := parse_expr(p, unary)

	return ast.PrefixExpression{
		Operator: operatorToken,
		Right:    expr,
	}
}

func parse_assignment_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	p.advance()
	rhs := parse_expr(p, bp)

	return ast.AssignmentExpression{
		Assigne:       left,
		AssignedValue: rhs,
	}
}

func parse_range_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	p.advance()

	return ast.RangeExpression{
		Lower: left,
		Upper: parse_expr(p, bp),
	}
}

func parse_binary_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	operatorToken := p.advance()
	right := parse_expr(p, defalt_bp)

	return ast.BinaryExpression{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parse_primary_expr(p *parser) ast.Expression {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpression{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpression{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpression{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Cannot create primary_expr from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parse_member_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	isComputed := p.advance().Kind == lexer.OPEN_BRACKET

	if isComputed {
		rhs := parse_expr(p, bp)
		p.expect(lexer.CLOSE_BRACKET)
		return ast.ComputedExpression{
			Member:   left,
			Property: rhs,
		}
	}

	return ast.MemberExpression{
		Member:   left,
		Property: p.expect(lexer.IDENTIFIER).Value,
	}
}

func parse_array_literal_expr(p *parser) ast.Expression {
	p.expect(lexer.OPEN_BRACKET)
	arrayContents := make([]ast.Expression, 0)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_BRACKET {
		arrayContents = append(arrayContents, parse_expr(p, logical))

		if !p.currentToken().IsOneOfMany(lexer.EOF, lexer.CLOSE_BRACKET) {
			p.expect(lexer.COMMA)
		}
	}

	p.expect(lexer.CLOSE_BRACKET)

	return ast.ArrayLiteral{
		Contents: arrayContents,
	}
}

func parse_grouping_expr(p *parser) ast.Expression {
	p.expect(lexer.OPEN_PAREN)
	expr := parse_expr(p, defalt_bp)
	p.expect(lexer.OPEN_PAREN)
	return expr
}

func parse_call_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	p.advance()
	arguments := make([]ast.Expression, 0)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_PAREN {
		arguments = append(arguments, parse_expr(p, assignment))

		if !p.currentToken().IsOneOfMany(lexer.EOF, lexer.CLOSE_PAREN) {
			p.expect(lexer.COMMA)
		}
	}

	p.expect(lexer.CLOSE_PAREN)
	return ast.CallExpression{
		Method:    left,
		Arguments: arguments,
	}
}

func parse_fn_expr(p *parser) ast.Expression {
	p.expect(lexer.FN)
	functionParams, returnType, functionBody := parse_fn_params_and_body(p)

	return ast.FunctionExpression{
		Parameters: functionParams,
		ReturnType: returnType,
		Body:       functionBody,
	}
}

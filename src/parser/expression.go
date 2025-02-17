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
		panic(fmt.Sprintf("NUD handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)
	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
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
		panic(fmt.Sprintf("Cannot create primary_expression from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parse_binary_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	operatorToken := p.advance()
	right := parse_expr(p, bp)

	return ast.BinaryExpression{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

package parser

import (
	"github.com/Kresqle/genuml/src/ast"
	"github.com/Kresqle/genuml/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	createTypeTokenLookups()

	p := &parser{
		tokens: tokens,
		pos:    0,
	}

	return p
}

func Parse(source string) ast.BlockStatement {
	tokens := lexer.Tokenize(source)
	p := createParser(tokens)
	body := make([]ast.Statement, 0)

	for p.hasTokens() {
		body = append(body, parse_stmt(p))
	}

	return ast.BlockStatement{
		Body: body,
	}
}

package main

import (
	"os"

	"github.com/Kresqle/genuml/src/lexer"
	"github.com/Kresqle/genuml/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/03.lang")
	tokens := lexer.Tokenize(string(bytes))
	ast := parser.Parse(tokens)
	litter.Dump(ast)
}

package parser

func Parse(f string) *AST {
	tokens := Lex(f)
	ast := new(AST)
	return ast
}

type AST struct{}

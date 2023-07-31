package parser

type TokenType uint64

const (
	CommentChar = '#'
)

var SingleCharacterTokens = map[rune]TokenType{
	'(': LEFT_PAREN,
	')': RIGHT_PAREN,
	'{': LEFT_BRACE,
	'}': RIGHT_BRACE,
	',': COMMA,
	'.': DOT,
	'-': MINUS,
	'+': PLUS,
	';': SEMICOLON,
	'/': SLASH,
	'*': STAR,
    '_': UNDERSCORE,
}

type oneOrTwoCharacterToken struct {
	Id       TokenType
	NextRune rune
	NextId   TokenType
}

var OneOrTwoCharacterTokens = map[rune]oneOrTwoCharacterToken{
	'!': {BANG, '=', BANG_EQUAL},
	'=': {EQUAL, '=', EQUAL_EQUAL},
	'>': {GREATER, '=', GREATER_EQUAL},
	'<': {LESS, '=', LESS_EQUAL},
}

var Keywords = map[string]TokenType{
	"and":   AND,
	"else":  ELSE,
	"false": FALSE,
	"for":   FOR,
	"fun":   FUN,
	"if":    IF,
	"let":   LET,
	"or":    OR,
	"true":  TRUE,
	"while": WHILE,
	"yank":  YANK,
	"yield": YIELD,
}

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
    UNDERSCORE

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	ELSE
	FALSE
	FUN
	FOR
	IF
	LET
	OR
	TRUE
	WHILE
	YANK
	YIELD

	EOF
)

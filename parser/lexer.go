package parser

import (
	"fmt"
	"strings"
)

func Lex(f string) *tokenList {
	tokens := tokenList{
		head: nil,
		tail: nil,
	}
	reader := Reader(f)
	for reader.Read() {
		if isWhitespace(reader.Current) {
			continue
		} else if reader.Current == CommentChar {
			for reader.Read() && reader.Current != '\n' {
			}
		} else if tokenType, ok := SingleCharacterTokens[reader.Current]; ok {
			tokens.Append(
				newToken(tokenType, string(reader.Current), "", reader.Line, reader.Col, &reader.LineContent, &f),
			)
		} else if tokenType, ok := OneOrTwoCharacterTokens[reader.Current]; ok {
			if reader.Peak(1) == tokenType.NextRune {
				lexeme := fmt.Sprintf("%c%c", reader.Current, reader.Peak(1))
				reader.Read()
				tokens.Append(newToken(tokenType.NextId, lexeme, "", reader.Line, reader.Col, &reader.LineContent, &f))
			} else {
				tokens.Append(newToken(tokenType.Id, string(reader.Current), "", reader.Line, reader.Col, &reader.LineContent, &f))
			}
		} else if isDigit(reader.Current) {
			stringBuilder := strings.Builder{}
			stringBuilder.WriteRune(reader.Current)
			for reader.Read() && isDigit(reader.Current) {
				stringBuilder.WriteRune(reader.Current)
			}
			if reader.Current == '.' {
				stringBuilder.WriteRune(reader.Current)
				for reader.Read() && isDigit(reader.Current) {
					stringBuilder.WriteRune(reader.Current)
				}
			}
			lexeme := stringBuilder.String()
			tokens.Append(
				newToken(
					NUMBER,
					lexeme,
					lexeme,
					reader.Line,
					reader.Col,
					&reader.LineContent,
					&f,
				),
			)
		} else if isAlpha(reader.Current) {
			stringBuilder := strings.Builder{}
			stringBuilder.WriteRune(reader.Current)
			for reader.Read() && isAlphaNumeric(reader.Current) {
				stringBuilder.WriteRune(reader.Current)
			}
			lexeme := stringBuilder.String()
			if tokenType, ok := Keywords[lexeme]; ok {
				tokens.Append(
					newToken(
						tokenType,
						lexeme,
						lexeme,
						reader.Line,
						reader.Col,
						&reader.LineContent,
						&f,
					),
				)
			} else {
				tokens.Append(
					newToken(
						IDENTIFIER,
						lexeme,
						lexeme,
						reader.Line,
						reader.Col,
						&reader.LineContent,
						&f,
					),
				)
			}
		} else if reader.Current == '"' {
			stringBuilder := strings.Builder{}
		string_literal_loop:
			for reader.Read() {
				switch reader.Current {
				case '\\':
					reader.Read()
					switch reader.Current {
					case '"', '\\':
						stringBuilder.WriteRune(reader.Current)
					case 'n':
						stringBuilder.WriteRune('\n')
					case 't':
						stringBuilder.WriteRune('\t')
					case 'r':
						stringBuilder.WriteRune('\r')
					default:
						lexingError("Invalid escape sequence", reader.Line, reader.Col, reader.LineContent, f)
					}
				case '"':
					tokens.Append(
						newToken(
							STRING,
							stringBuilder.String(),
							stringBuilder.String(),
							reader.Line,
							reader.Col,
							&reader.LineContent,
							&f,
						),
					)
					break string_literal_loop
				default:
					stringBuilder.WriteRune(reader.Current)
				}
			}
			if reader.Current != '"' {
				lexingError("Unterminated string", reader.Line, reader.Col, reader.LineContent, f)
			}
		} else {
			lexingError(fmt.Sprintf("Unexpected character '%c'", reader.Current), reader.Line, reader.Col, reader.LineContent, f)
		}
	}
	tokens.Append(
		newToken(
			EOF,
			"",
			"",
			reader.Line,
			reader.Col,
			&reader.LineContent,
			&f,
		),
	)
	return &tokens
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isDigit(r)
}

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		r == '_'
}

func isDigit(r rune) bool {
	return (r >= '0' && r <= '9') || r == '_' || r == '.'
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n' || r == '\v' ||
		r == '\f'
}

type Token struct {
	Type        TokenType
	Lexeme      string
	Literal     string
	Line        uint64
	Col         uint64
	LineContent *string
	File        *string
}

func newToken(
	tokenType TokenType,
	lexeme string,
	literal string,
	line uint64,
	col uint64,
	lineContent *string,
	file *string,
) *Token {
	return &Token{
		Type:        tokenType,
		Lexeme:      lexeme,
		Literal:     literal,
		Line:        line,
		Col:         col,
		LineContent: lineContent,
		File:        file,
	}
}

func lexingError(
	message string,
	line uint64,
	col uint64,
	text string,
	file string,
) {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("Lexing error: ")
	stringBuilder.WriteString(message)
	stringBuilder.WriteString(fmt.Sprintf("\n%s:%d:%d\n", file, line, col))
	stringBuilder.WriteString(text)
	stringBuilder.WriteRune('\n')
	for i := uint64(2); i < col; i++ {
		stringBuilder.WriteRune(' ')
	}
	stringBuilder.WriteRune('^')
	fmt.Println(stringBuilder.String())
}

package lexer

import (
	"strings"
)

// TokenType — тип токенів
type TokenType string

const (
	TOKEN_HAI     TokenType = "HAI"
	TOKEN_VISIBLE TokenType = "VISIBLE"
	TOKEN_KTHXBYE TokenType = "KTHXBYE"
	TOKEN_UNKNOWN TokenType = "UNKNOWN"
)

// Token — структура токена
type Token struct {
	Type    TokenType
	Literal string
}

// Tokenize — базовий лексичний аналізатор
func Tokenize(input string) []Token {
	words := strings.Fields(input)
	var tokens []Token

	for _, word := range words {
		switch word {
		case "HAI":
			tokens = append(tokens, Token{Type: TOKEN_HAI, Literal: word})
		case "VISIBLE":
			tokens = append(tokens, Token{Type: TOKEN_VISIBLE, Literal: word})
		case "KTHXBYE":
			tokens = append(tokens, Token{Type: TOKEN_KTHXBYE, Literal: word})
		default:
			tokens = append(tokens, Token{Type: TOKEN_UNKNOWN, Literal: word})
		}
	}

	return tokens
}

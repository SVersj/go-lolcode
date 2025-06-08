package parser

import (
	"errors"
	"github.com/SVersj/go-lolcode/internal/lexer"
)

type Statement struct {
	Command string
	Args    []string
}

func Parse(tokens []lexer.Token) ([]Statement, error) {
	var statements []Statement

	if len(tokens) < 2 || tokens[0].Type != lexer.TOKEN_HAI || tokens[len(tokens)-1].Type != lexer.TOKEN_KTHXBYE {
		return nil, errors.New("program must start with HAI and end with KTHXBYE")
	}

	for i := 1; i < len(tokens)-1; i++ {
		token := tokens[i]

		if token.Type == lexer.TOKEN_VISIBLE {
			if i+1 < len(tokens) {
				arg := tokens[i+1].Literal
				statements = append(statements, Statement{Command: "VISIBLE", Args: []string{arg}})
				i++ // Пропускаємо аргумент
			} else {
				return nil, errors.New("VISIBLE without argument")
			}
		}
	}

	return statements, nil
}

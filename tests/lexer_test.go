package tests

import (
	"testing"

	"github.com/SVersj/go-lolcode/internal/lexer"
)

func TestTokenize(t *testing.T) {
	input := "HAI VISIBLE Hello KTHXBYE"
	expected := []lexer.Token{
		{Type: lexer.TOKEN_HAI, Literal: "HAI"},
		{Type: lexer.TOKEN_VISIBLE, Literal: "VISIBLE"},
		{Type: lexer.TOKEN_UNKNOWN, Literal: "Hello"},
		{Type: lexer.TOKEN_KTHXBYE, Literal: "KTHXBYE"},
	}

	tokens := lexer.Tokenize(input)

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, tok := range tokens {
		if tok != expected[i] {
			t.Errorf("token %d: expected %+v, got %+v", i, expected[i], tok)
		}
	}
}

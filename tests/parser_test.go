package tests

import (
	"testing"

	"github.com/SVersj/go-lolcode/internal/lexer"
	"github.com/SVersj/go-lolcode/internal/parser"
)

func TestParse(t *testing.T) {
	tokens := lexer.Tokenize("HAI VISIBLE Hello KTHXBYE")

	statements, err := parser.Parse(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(statements))
	}

	stmt := statements[0]
	if stmt.Command != "VISIBLE" {
		t.Errorf("expected VISIBLE command, got %s", stmt.Command)
	}
	if len(stmt.Args) != 1 || stmt.Args[0] != "Hello" {
		t.Errorf("expected argument 'Hello', got %v", stmt.Args)
	}
}

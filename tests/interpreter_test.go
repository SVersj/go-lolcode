package tests

import (
	"bytes"
	"github.com/SVersj/go-lolcode/internal/interpreter"
	"github.com/SVersj/go-lolcode/internal/parser"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	statements := []parser.Statement{
		{Command: "VISIBLE", Args: []string{"Hello", "World"}},
	}

	// Перехоплюємо stdout
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	_ = interpreter.Run(statements)

	w.Close()
	os.Stdout = stdout
	_, _ = buf.ReadFrom(r)

	output := buf.String()
	expected := "Hello\nWorld\n"

	if strings.TrimSpace(output) != strings.TrimSpace(expected) {
		t.Errorf("expected output:\n%s\ngot:\n%s", expected, output)
	}
}

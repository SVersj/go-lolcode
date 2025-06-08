package interpreter

import (
	"fmt"
	"github.com/SVersj/go-lolcode/internal/parser"
)

func Run(statements []parser.Statement) error {
	for _, stmt := range statements {
		switch stmt.Command {
		case "VISIBLE":
			for _, arg := range stmt.Args {
				fmt.Println(arg)
			}
		default:
			return fmt.Errorf("unsupported command: %s", stmt.Command)
		}
	}
	return nil
}

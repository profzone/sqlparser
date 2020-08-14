package sql

import (
	"fmt"
	"github.com/profzone/sqlparser/pkg/lexer"
	"testing"
)

func init() {
	lexer.RegisterTokenFunc("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_", scanIdent)
	lexer.RegisterTokenFunc("\"", scanString)
}

func TestScan(t *testing.T) {
	scanner := lexer.NewScanner("\"1231123\" 123 456")
	for !scanner.Eof() {
		tok, pos, lit := scanner.Scan()
		fmt.Println(tok, pos, lit)
	}
}

package calc

import (
	"github.com/profzone/sqlparser/pkg/lexer"
	"unicode"
)

func init() {
	lexer.RegisterTokenFunc("0123456789", scanInteger)

	lexer.RegisterToken('+', ADD)
	lexer.RegisterToken('-', MIN)
	lexer.RegisterToken('*', SUB)
	lexer.RegisterToken('/', DIV)
	lexer.RegisterToken('%', MOD)
	lexer.RegisterToken('(', LP)
	lexer.RegisterToken(')', RP)

	lexer.RegisterTokenStr("add", ADD)
	lexer.RegisterTokenStr("min", MIN)
	lexer.RegisterTokenStr("sub", SUB)
	lexer.RegisterTokenStr("div", DIV)
	lexer.RegisterTokenStr("mod", MOD)
}

func scanInteger(s *lexer.Scanner) (int, lexer.Pos, string) {
	pos := s.Pos()
	s.Inc()
	s.IncAsLongAs(unicode.IsDigit)
	return INTEGER, pos, s.Data(&pos)
}

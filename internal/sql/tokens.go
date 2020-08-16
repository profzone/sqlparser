package sql

import (
	"github.com/profzone/sqlparser/internal/common"
	"github.com/profzone/sqlparser/pkg/lexer"
	"strings"
)

func init() {
	//lexer.RegisterTokenFunc("0123456789", scanInteger)

	lexer.RegisterToken('+', OPERATOR)
	lexer.RegisterToken('-', OPERATOR)
	lexer.RegisterToken('*', OPERATOR) // TODO
	lexer.RegisterToken('/', OPERATOR)
	lexer.RegisterToken('%', OPERATOR)
	lexer.RegisterToken('(', LP)
	lexer.RegisterToken(')', RP)
	lexer.RegisterToken('.', DOT)
	lexer.RegisterToken(',', COMMA)
	lexer.RegisterToken('*', STAR) // TODO

	lexer.RegisterTokenStr(">", LOGICAL)
	lexer.RegisterTokenStr(">=", LOGICAL)
	lexer.RegisterTokenStr("<", LOGICAL)
	lexer.RegisterTokenStr("<=", LOGICAL)
	lexer.RegisterTokenStr("<>", LOGICAL)
	lexer.RegisterTokenStr("!=", LOGICAL)
	lexer.RegisterTokenStr("=", LOGICAL)

	lexer.RegisterTokenFunc("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_", scanIdent)
	lexer.RegisterTokenFunc("\"", scanString)
}

var tokenMap = map[string]int{
	"SELECT": SELECT,
	"AS":     AS,
	"FROM":   FROM,
	"WHERE":  WHERE,
	"ORDER":  ORDER,
	"GROUP":  GROUP,
	"BY":     BY,
	"DESC":   DESC,
	"ASC":    ASC,
}

var logicalMap = map[string]int{
	"AND": LOGICAL,
	"OR":  LOGICAL,
}

//func scanInteger(s *lexer.Scanner) (int, lexer.Pos, string) {
//	pos := s.Pos()
//	s.Inc()
//	s.IncAsLongAs(unicode.IsDigit)
//	return IDENT, pos, s.Data(&pos)
//}

func scanIdent(s *lexer.Scanner) (int, lexer.Pos, string) {
	pos := s.Pos()
	s.Inc()
	s.IncAsLongAs(common.IsIdentChar)
	return IDENT, pos, s.Data(&pos)
}

func scanString(s *lexer.Scanner) (int, lexer.Pos, string) {
	pos := s.Pos()
	for !s.Eof() {
		s.Inc()
		ch := s.Peek()
		if ch == '"' {
			s.Inc()
			break
		}
	}

	return CSTRING, pos, strings.Trim(s.Data(&pos), "\"")
}

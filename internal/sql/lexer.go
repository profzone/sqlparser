package sql

import (
	"github.com/profzone/sqlparser/internal/sql/ast"
	"github.com/profzone/sqlparser/pkg/lexer"
	"strconv"
	"strings"
)

type Lexer struct {
	*lexer.Scanner
	Root ast.Node
	errs []string
}

func NewLexer(code string) *Lexer {
	return &Lexer{
		Scanner: lexer.NewScanner(code),
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, _, lit := l.Scan()
	switch tok {
	case IDENT:
		// if is keywords
		if t, ok := tokenMap[strings.ToUpper(lit)]; ok {
			return t
		}
		// if is logical
		if t, ok := logicalMap[strings.ToUpper(lit)]; ok {
			lval.logical = lit
			return t
		}
		// if is integer constant
		i, err := strconv.ParseInt(lit, 10, 64)
		if err == nil {
			lval.int64 = i
			return CINTEGER
		}
		// if is float constant
		f, err := strconv.ParseFloat(lit, 64)
		if err == nil {
			lval.float64 = f
			return CFLOAT
		}
		// if is bool constant
		b, err := strconv.ParseBool(lit)
		if err == nil {
			lval.bool = b
			return CBOOL
		}
		lval.string = lit
	case CSTRING:
		lval.string = lit
	case LOGICAL:
		lval.logical = lit
	case OPERATOR:
		lval.operator = lit
	}

	return tok
}

func (l *Lexer) Error(s string) {
	l.errs = append(l.errs, s)
}

func (l *Lexer) Errors() []string {
	return l.errs
}

func (l *Lexer) Reset(code string) {
	l.Root = nil
	l.Scanner.Reset(code)
	l.errs = []string{}
}

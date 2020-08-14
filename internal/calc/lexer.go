package calc

import (
	"github.com/profzone/sqlparser/internal/calc/ast"
	"github.com/profzone/sqlparser/pkg/lexer"
	"strconv"
)

type Lexer struct {
	*lexer.Scanner
	RootNode ast.ExprNode
	errs     []string
}

func NewLexer(code string) *Lexer {
	return &Lexer{
		Scanner: lexer.NewScanner(code),
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, _, lit := l.Scan()
	switch tok {
	case INTEGER:
		i, err := strconv.ParseInt(lit, 10, 64)
		if err != nil {
			return 0
		}
		lval.int64 = i
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
	l.RootNode = nil
	l.Scanner.Reset(code)
}

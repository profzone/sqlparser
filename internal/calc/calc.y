%{
package calc

import (
	"github.com/profzone/sqlparser/src/calc/ast"
)
%}

%union {
	int64 int64
	expr ast.ExprNode
	str string
}

%token <str> ADD MIN SUB DIV MOD LP RP
%token <int64> INTEGER

%left ADD MIN SUB DIV MOD

%type <expr> Expr ConstExpr /*AddExpr MinExpr SubExpr DivExpr ModExpr*/

%%

program:
	Expr
	| '\n'

Expr:
	LP Expr RP
	{
		expr := &ast.PrecedenceExpr{}
		expr.Val = $2
		$$ = expr
	}
	| Expr ADD Expr
	{
		expr := &ast.AddExpr{}
		expr.Left = $1
		expr.Right = $3
		expr.Op = ADD
		$$ = expr
		(yylex.(*Lexer)).RootNode = $$
	}
	| Expr MIN Expr
	{
		expr := &ast.MinExpr{}
		expr.Left = $1
		expr.Right = $3
		expr.Op = MIN
		$$ = expr
		(yylex.(*Lexer)).RootNode = $$
	}
	| Expr SUB Expr
	{
		expr := &ast.SubExpr{}
		expr.Left = $1
		expr.Right = $3
		expr.Op = SUB
		$$ = expr
		(yylex.(*Lexer)).RootNode = $$
	}
	| Expr DIV Expr
	{
		expr := &ast.DivExpr{}
		expr.Left = $1
		expr.Right = $3
		expr.Op = DIV
		$$ = expr
		(yylex.(*Lexer)).RootNode = $$
	}
	| Expr MOD Expr
	{
		expr := &ast.ModExpr{}
		expr.Left = $1
		expr.Right = $3
		expr.Op = MOD
		$$ = expr
		(yylex.(*Lexer)).RootNode = $$
	}
	| ConstExpr
	;

ConstExpr:
	INTEGER
	{
		expr := &ast.ConstExpr{}
		expr.Val = $1
		$$ = expr
	}

%%

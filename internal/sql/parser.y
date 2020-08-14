%{
package sql

import (
	"github.com/profzone/sqlparser/internal/sql/ast"
)
%}

%union {
	string string
	int64 int64
	float64 float64
	bool bool

	operator string
	logical string
	node ast.Node
	expr ast.Expr
	stmt ast.Statement
	ident *ast.Ident
	selector *ast.SelectorExpr

	fields ast.FieldsNode
	from *ast.FromNode
	twoOpExpr *ast.TwoOpExpr
	conditions ast.TwoOpExprs
	where *ast.WhereNode
	having *ast.HavingNode
}

// operators
%token <string> LP RP COMMA DOT
%token <operator> OPERATOR
%token <logical> LOGICAL
// keywords
%token <string> SELECT STAR AS FROM WHERE GROUP ORDER HAVING LIMIT OFFSET
// identities
%token <string> IDENT
// constants
%token <string> CSTRING
%token <int64> CINTEGER
%token <float64> CFLOAT
%token <bool> CBOOL

%left OPERATOR LOGICAL COMMA DOT

%type <logical> Logical
%type <operator> Operator
%type <expr> Expr Exprs Conditions Const
%type <ident> Ident
%type <selector> Selector
%type <node> Field
%type <twoOpExpr> Condition
%type <fields> FieldNode
%type <from> FromNode
%type <where> WhereNode
%type <stmt> SelectStatement

%%

Program:
	SelectStatement

Logical:
	LOGICAL
	{
		$$ = $1
	}

Operator:
	OPERATOR
	{
		$$ = $1
	}

Ident:
	IDENT
	{
		node := &ast.Ident{}
		node.SetVal($1)
		$$ = node
	}

Const:
	CSTRING
	{
		node := &ast.ConstExpr{
			Val: $1,
		}
		$$ = node
	}
	| CINTEGER
	{
		node := &ast.ConstExpr{
			Val: $1,
		}
		$$ = node
	}
	| CFLOAT
	{
		node := &ast.ConstExpr{
			Val: $1,
		}
		$$ = node
	}
	| CBOOL
	{
		node := &ast.ConstExpr{
			Val: $1,
		}
		$$ = node
	}

Selector:
	Ident
	{
		node := &ast.SelectorExpr{
			X: $1,
		}
		$$ = node
	}
	| Selector DOT Ident
	{
		node := &ast.SelectorExpr{
			Sel: $1,
			X: $3,
		}
		$$ = node
	}

Field:
	STAR
	{
		node := &ast.Field{
			All: true,
		}
		$$ = node
	}
	| Selector
	{
		node := &ast.Field{
			Name: $1,
		}
		$$ = node
	}
	| Selector AS Ident
	{
		node := &ast.Field{
			Name: $1,
			Alias: $3,
		}
		$$ = node
	}

FieldNode:
	Field
	{
		node := ast.FieldsNode{}
		node.Fields = append(node.Fields, *($1.(*ast.Field)))
		$$ = node
	}
	| Field COMMA FieldNode
	{
		$$.Fields = append($3.Fields, *($1.(*ast.Field)))
	}

FromNode:
	FROM Selector
	{
		node := &ast.FromNode{
			Name: $2,
		}
		$$ = node
	}
	| FROM Selector AS Ident
	{
		node := &ast.FromNode{
			Name: $2,
			Alias: $4,
		}
		$$ = node
	}

Expr:
	Selector
	{
		$$ = $1
	}
	| Const
	{
		$$ = $1
	}

Exprs:
	Expr
	{
		$$ = $1
	}
	| Exprs Operator Expr
	{
		node := &ast.TwoOpExpr{
			Left: $1,
			Op: $2,
			Right: $3,
		}
		$$ = node
	}

Condition:
	Exprs Logical Exprs
	{
		node := &ast.TwoOpExpr{
			Left: $1,
			Op: $2,
			Right: $3,
		}
		$$ = node
	}

Conditions:
	Condition
	{
		$$ = $1
	}
	| LP Conditions RP
	{
		node := &ast.PrecedenceExpr{
			Val: $2,
		}
		$$ = node
	}
	| Conditions Logical Condition
	{
		node := &ast.TwoOpExpr{
			Left: $1,
			Op: $2,
			Right: $3,
		}
		$$ = node
	}

WhereNode:
	WHERE Conditions
	{
		node := &ast.WhereNode{
			Conditions: $2,
		}
		$$ = node
	}

SelectStatement:
	SELECT FieldNode
	{
		stmt := &ast.SelectStatement{
			Fields: $2,
		}
		$$ = stmt
		(yylex.(*Lexer)).Root = $$
	}
	| SELECT FieldNode FromNode
	{
		stmt := &ast.SelectStatement{
			Fields: $2,
			From: $3,
		}
		$$ = stmt
		(yylex.(*Lexer)).Root = $$
	}
	| SELECT FieldNode FromNode WhereNode
	{
		stmt := &ast.SelectStatement{
			Fields: $2,
			From: $3,
			Where: $4,
		}
		$$ = stmt
		(yylex.(*Lexer)).Root = $$
	}

%%
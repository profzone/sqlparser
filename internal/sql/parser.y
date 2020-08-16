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
	orders *ast.OrderNode
	groups *ast.GroupNode
	orderList []ast.Order
	groupList []ast.Group
}

// operators
%token <string> LP RP COMMA DOT
%token <operator> OPERATOR
%token <logical> LOGICAL
// keywords
%token <string> SELECT STAR AS FROM WHERE GROUP ORDER BY HAVING LIMIT OFFSET DESC ASC
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
%type <node> Field Order Group
%type <twoOpExpr> Condition
%type <fields> FieldNode
%type <from> FromNode
%type <where> WhereNode
%type <orderList> Orders
%type <orders> OrderNode
%type <groupList> Groups
%type <groups> GroupNode
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
	{
		$$ = nil
	}
	| WHERE Conditions
	{
		node := &ast.WhereNode{
			Conditions: $2,
		}
		$$ = node
	}

Group:
	Expr
	{
		node := &ast.Group{
			Expr: $1,
		}
		$$ = node
	}

Groups:
	Group
	{
		$$ = []ast.Group{*($1.(*ast.Group))}
	}
	| Groups COMMA Group
	{
		$$ = append($1, *($3.(*ast.Group)))
	}

GroupNode:
	{
		$$ = nil
	}
	| GROUP BY Groups
	{
		node := &ast.GroupNode{
			Groups: $3,
		}
		$$ = node
	}

Order:
	Expr
	{
		node := &ast.Order{
			Expr: $1,
		}
		$$ = node
	}
	| Expr DESC
	{
		node := &ast.Order{
			Expr: $1,
			Order: 1,
		}
		$$ = node
	}
	| Expr ASC
	{
		node := &ast.Order{
			Expr: $1,
		}
		$$ = node
	}

Orders:
	Order
	{
		$$ = []ast.Order{*($1.(*ast.Order))}
	}
	| Orders COMMA Order
	{
		$$ = append($1, *($3.(*ast.Order)))
	}

OrderNode:
	{
		$$ = nil
	}
	| ORDER BY Orders
	{
		node := &ast.OrderNode{
			Orders: $3,
		}
		$$ = node
	}

// select id, age, name from t_user where id=1 group by id having id=1 order by id desc
SelectStatement:
	SELECT FieldNode
	{
		stmt := &ast.SelectStatement{
			Fields: $2,
		}
		$$ = stmt
		(yylex.(*Lexer)).Root = $$
	}
	| SELECT FieldNode FromNode WhereNode GroupNode OrderNode
	{
		stmt := &ast.SelectStatement{
			Fields: $2,
			From: $3,
			Where: $4,
			Group: $5,
			Order: $6,
		}
		$$ = stmt
		(yylex.(*Lexer)).Root = $$
	}

%%

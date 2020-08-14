package ast

type exprNode struct{}

var (
	_ ExprNode = &AddExpr{}
	_ ExprNode = &MinExpr{}
	_ ExprNode = &SubExpr{}
	_ ExprNode = &DivExpr{}
	_ ExprNode = &ModExpr{}

	_ ExprNode = &PrecedenceExpr{}
)

type ConstExpr struct {
	exprNode
	Val interface{}
}

func (c *ConstExpr) Visit(visitor Visitor) bool {
	visitor.Enter(c)
	return visitor.Leave(c)
}

func (c *ConstExpr) GetVal() interface{} {
	return c.Val
}

type TwoOpExpr struct {
	exprNode
	Op    int
	Left  ExprNode
	Right ExprNode
}

func (expr *TwoOpExpr) Visit(visitor Visitor) bool {
	if !visitor.Enter(expr) {
		return visitor.Leave(expr)
	}

	if !expr.Left.Visit(visitor) {
		return false
	}
	if !expr.Right.Visit(visitor) {
		return false
	}

	return visitor.Leave(expr)
}

func (expr *TwoOpExpr) GetVal() interface{} {
	return nil
}

type AddExpr struct {
	TwoOpExpr
}

func (expr *AddExpr) GetVal() interface{} {
	return nil
}

type MinExpr struct {
	TwoOpExpr
}

func (expr *MinExpr) GetVal() interface{} {
	return nil
}

type SubExpr struct {
	TwoOpExpr
}

func (expr *SubExpr) GetVal() interface{} {
	return nil
}

type DivExpr struct {
	TwoOpExpr
}

func (expr *DivExpr) GetVal() interface{} {
	return nil
}

type ModExpr struct {
	TwoOpExpr
}

func (expr *ModExpr) GetVal() interface{} {
	return nil
}

type PrecedenceExpr struct {
	exprNode
	Val ExprNode
}

func (p *PrecedenceExpr) Visit(visitor Visitor) bool {
	if !visitor.Enter(p) {
		return visitor.Leave(p)
	}
	if p.Val.Visit(visitor) {
		return false
	}
	return visitor.Leave(p)
}

func (p *PrecedenceExpr) GetVal() interface{} {
	return p.Val.GetVal()
}

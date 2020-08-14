package ast

type Field struct {
	Name  Expr
	Func  Expr // TODO
	Alias Expr
	Val   *ConstExpr // TODO
	All   bool
}

func (f *Field) Visit(visitor Visitor) bool {
	if !visitor.Enter(f) {
		return visitor.Leave(f)
	}

	if f.Name != nil && !f.Name.Visit(visitor) {
		return false
	}
	if f.Func != nil && !f.Func.Visit(visitor) {
		return false
	}
	if f.Alias != nil && !f.Alias.Visit(visitor) {
		return false
	}
	if f.Val != nil && !f.Val.Visit(visitor) {
		return false
	}

	return visitor.Leave(f)
}

type FieldsNode struct {
	Fields []Field
}

func (f FieldsNode) Visit(visitor Visitor) bool {
	if !visitor.Enter(f) {
		return visitor.Leave(f)
	}

	for _, field := range f.Fields {
		if !field.Visit(visitor) {
			return false
		}
	}

	return visitor.Leave(f)
}

type FromNode struct {
	Name         *SelectorExpr
	SubStatement Statement
	Alias        Expr
}

func (f *FromNode) Visit(visitor Visitor) bool {
	if !visitor.Enter(f) {
		return visitor.Leave(f)
	}

	if f.Name != nil && !f.Name.Visit(visitor) {
		return false
	}
	if f.SubStatement != nil && !f.SubStatement.Visit(visitor) {
		return false
	}
	if f.Alias != nil && !f.Alias.Visit(visitor) {
		return false
	}

	return visitor.Leave(f)
}

type WhereNode struct {
	Conditions Expr
}

func (w *WhereNode) Visit(visitor Visitor) bool {
	if !visitor.Enter(w) {
		return visitor.Leave(w)
	}

	if !w.Conditions.Visit(visitor) {
		return false
	}

	return visitor.Leave(w)
}

type HavingNode struct {
	Conditions Expr
}

func (w *HavingNode) Visit(visitor Visitor) bool {
	if !visitor.Enter(w) {
		return visitor.Leave(w)
	}

	if !w.Conditions.Visit(visitor) {
		return false
	}

	return visitor.Leave(w)
}

type Order struct {
	Name  *SelectorExpr
	Expr  Expr // TODO
	Order int
}

func (o *Order) Visit(visitor Visitor) bool {
	if !visitor.Enter(o) {
		return visitor.Leave(o)
	}

	if o.Name != nil && !o.Name.Visit(visitor) {
		return false
	}
	if o.Expr != nil && !o.Expr.Visit(visitor) {
		return false
	}

	return visitor.Leave(o)
}

type OrderNode struct {
	Orders []Order
}

func (o *OrderNode) Visit(visitor Visitor) bool {
	if !visitor.Enter(o) {
		return visitor.Leave(o)
	}

	for _, c := range o.Orders {
		if !c.Visit(visitor) {
			return false
		}
	}

	return visitor.Leave(o)
}

type Group struct {
	Name *SelectorExpr
	Expr Expr // TODO
}

func (g *Group) Visit(visitor Visitor) bool {
	if !visitor.Enter(g) {
		return visitor.Leave(g)
	}

	if g.Name != nil && !g.Name.Visit(visitor) {
		return false
	}
	if g.Expr != nil && !g.Expr.Visit(visitor) {
		return false
	}

	return visitor.Leave(g)
}

type GroupNode struct {
	Groups []Group
}

func (g *GroupNode) Visit(visitor Visitor) bool {
	if !visitor.Enter(g) {
		return visitor.Leave(g)
	}

	for _, c := range g.Groups {
		if !c.Visit(visitor) {
			return false
		}
	}

	return visitor.Leave(g)
}

type LimitNode struct {
	Limit  int
	Offset int
}

func (l *LimitNode) Visit(visitor Visitor) bool {
	visitor.Enter(l)
	return visitor.Leave(l)
}

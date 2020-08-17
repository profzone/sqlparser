package ast

type Field struct {
	Expr  Expr
	Alias Expr
	All   bool
}

func (f *Field) Visit(visitor Visitor) bool {
	if !visitor.Enter(f) {
		return visitor.Leave(f)
	}

	if f.Expr != nil && !f.Expr.Visit(visitor) {
		return false
	}
	if f.Alias != nil && !f.Alias.Visit(visitor) {
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
	Expr  Expr
	Alias Expr
}

func (f *FromNode) Visit(visitor Visitor) bool {
	if !visitor.Enter(f) {
		return visitor.Leave(f)
	}

	if f.Expr != nil && !f.Expr.Visit(visitor) {
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
	Expr  Expr // TODO
	Order int
}

func (o *Order) Visit(visitor Visitor) bool {
	if !visitor.Enter(o) {
		return visitor.Leave(o)
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
	Expr Expr
}

func (g *Group) Visit(visitor Visitor) bool {
	if !visitor.Enter(g) {
		return visitor.Leave(g)
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

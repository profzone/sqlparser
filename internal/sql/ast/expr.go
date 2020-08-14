package ast

import (
	"fmt"
	"strings"
)

type Visitor interface {
	Enter(n Node) bool
	Leave(n Node) bool
}

type Node interface {
	Visit(visitor Visitor) bool
}

type Expr interface {
	Node
	GetVal() interface{}
	SetVal(interface{}) error
}

type Ident struct {
	Name string
}

func (i *Ident) Visit(visitor Visitor) bool {
	visitor.Enter(i)
	return visitor.Leave(i)
}

func (i *Ident) GetVal() interface{} {
	return i.Name
}

func (i *Ident) SetVal(v interface{}) error {
	i.Name = fmt.Sprintf("%v", v)
	return nil
}

type ConstExpr struct {
	Val interface{}
}

func (i *ConstExpr) GetVal() interface{} {
	return i.Val
}

func (i *ConstExpr) SetVal(v interface{}) error {
	i.Val = v
	return nil
}

func (c *ConstExpr) Visit(visitor Visitor) bool {
	visitor.Enter(c)
	return visitor.Leave(c)
}

type SelectorExpr struct {
	Sel *SelectorExpr
	X   *Ident
}

func (s *SelectorExpr) Visit(visitor Visitor) bool {
	if !visitor.Enter(s) {
		return visitor.Leave(s)
	}

	if s.Sel != nil && !s.Sel.Visit(visitor) {
		return false
	}
	if s.X != nil && !s.X.Visit(visitor) {
		return false
	}

	return visitor.Leave(s)
}

func (s *SelectorExpr) GetVal() interface{} {
	return fmt.Sprintf("%s.%s", s.Sel.GetVal(), s.X.GetVal())
}

func (s *SelectorExpr) SetVal(i interface{}) error {
	var val = fmt.Sprintf("%v", i)
	var lastDot = strings.LastIndex(val, ".")
	if lastDot < 0 {
		s.Sel = nil
		s.X = &Ident{}
		return s.X.SetVal(val)
	} else {
		s.Sel = &SelectorExpr{}
		s.X = &Ident{}
		if err := s.Sel.SetVal(val[:lastDot]); err != nil {
			return err
		}
		return s.X.SetVal(val[lastDot+1:])
	}
}

type TwoOpExpr struct {
	Op    string
	Left  Expr
	Right Expr
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

func (expr *TwoOpExpr) SetVal(v interface{}) error {
	return nil
}

type TwoOpExprs []TwoOpExpr

func (t TwoOpExprs) Visit(visitor Visitor) bool {
	if !visitor.Enter(t) {
		return visitor.Leave(t)
	}
	for _, e := range t {
		if !e.Visit(visitor) {
			return false
		}
	}
	return visitor.Leave(t)
}

func (t TwoOpExprs) GetVal() interface{} {
	return nil
}

func (t TwoOpExprs) SetVal(i interface{}) error {
	return nil
}

type PrecedenceExpr struct {
	Val Expr
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

func (p *PrecedenceExpr) SetVal(v interface{}) error {
	return p.SetVal(v)
}

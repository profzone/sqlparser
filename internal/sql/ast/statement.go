package ast

import "fmt"

type Statement interface {
	Expr
	fmt.Stringer
}

type SelectStatement struct {
	Fields FieldsNode
	From   *FromNode
	Where  *WhereNode
	Group  *GroupNode
	Order  *OrderNode
	Having *HavingNode
	Limit  *LimitNode
}

func (s *SelectStatement) GetVal() interface{} {
	return nil
}

func (s *SelectStatement) SetVal(i interface{}) error {
	return nil
}

func (s *SelectStatement) Visit(visitor Visitor) bool {
	if !visitor.Enter(s) {
		return visitor.Leave(s)
	}

	if !s.Fields.Visit(visitor) {
		return false
	}
	if s.From != nil && !s.From.Visit(visitor) {
		return false
	}
	if s.Where != nil && !s.Where.Visit(visitor) {
		return false
	}
	if s.Group != nil && !s.Group.Visit(visitor) {
		return false
	}
	if s.Order != nil && !s.Order.Visit(visitor) {
		return false
	}
	if s.Having != nil && !s.Having.Visit(visitor) {
		return false
	}
	if s.Limit != nil && !s.Limit.Visit(visitor) {
		return false
	}

	return visitor.Leave(s)
}

func (s *SelectStatement) String() string {
	return ""
}

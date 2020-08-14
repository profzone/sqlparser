package ast

type Visitor interface {
	Enter(n Node) bool
	Leave(n Node) bool
}

type Node interface {
	Visit(visitor Visitor) bool
}

type ExprNode interface {
	Node
	GetVal() interface{}
}

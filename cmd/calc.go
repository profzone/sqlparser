package main

//
//import (
//	"bufio"
//	"fmt"
//	"github.com/profzone/sqlparser/internal/calc"
//	"github.com/profzone/sqlparser/internal/calc/ast"
//	"os"
//	"strings"
//)
//
//func main() {
//	inputReader := bufio.NewReader(os.Stdin)
//	lex := calc.NewLexer("")
//	visitor := &visitor{}
//	for {
//		fmt.Println("please input your code:")
//		str, _, err := inputReader.ReadLine()
//		if err != nil {
//			panic(err)
//		}
//		lex.Reset(strings.TrimSpace(string(str)))
//		calc.Parse(lex)
//
//		if errs := lex.Errors(); len(errs) > 0 {
//			fmt.Println(errs)
//			continue
//		}
//		lex.RootNode.Visit(visitor)
//	}
//}
//
//type visitor struct {}
//
//func (v *visitor) Enter(n ast.Node) bool {
//	switch val := n.(type) {
//	case *ast.TwoOpExpr:
//		switch val.Op {
//		case calc.ADD:
//			fmt.Println("AddExpr")
//		case calc.MIN:
//			fmt.Println("MinExpr")
//		case calc.SUB:
//			fmt.Println("SubExpr")
//		case calc.DIV:
//			fmt.Println("DivExpr")
//		case calc.MOD:
//			fmt.Println("ModExpr")
//		}
//	case *ast.ConstExpr:
//		fmt.Printf("ConstExpr: %v\n", val.GetVal())
//	}
//
//	return true
//}
//
//func (v *visitor) Leave(n ast.Node) bool {
//	return true
//}

package main

import (
	"bufio"
	"fmt"
	"github.com/profzone/sqlparser/internal/sql"
	"github.com/profzone/sqlparser/internal/sql/ast"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	lex := sql.NewLexer("")
	visitor := &visitor{}
	for {
		fmt.Println("please input your code:")
		str, _, err := inputReader.ReadLine()
		if err != nil {
			panic(err)
		}
		lex.Reset(strings.TrimSpace(string(str)))
		sql.Parse(lex)

		if errs := lex.Errors(); len(errs) > 0 {
			fmt.Println(errs)
			continue
		}

		lex.Root.Visit(visitor)
	}
}

type visitor struct{}

func (v *visitor) Enter(n ast.Node) bool {
	//switch val := n.(type) {
	//case *ast.SelectorExpr:
	//	fmt.Printf("SelectorExpr\n")
	//case *ast.Ident:
	//	fmt.Printf("Ident: %s\n", val.Name)
	//}

	return true
}

func (v *visitor) Leave(n ast.Node) bool {
	return true
}

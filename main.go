package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	file, err := parser.ParseFile(token.NewFileSet(), "doc.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(file.Doc.Text())

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			fmt.Println(strings.TrimSpace(d.Doc.Text()), decl.Pos())
		}
	}

	//for i, comment := range file.Comments {
	//	fmt.Println(strings.TrimSpace(comment.Text()), i)
	//}
}

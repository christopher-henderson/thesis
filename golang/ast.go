// package main

// import (
// 	"fmt"
// 	"go/ast"
// 	"go/parser"
// 	"go/token"
// 	"log"
// )

// func main() {
// 	fset := token.NewFileSet() // positions are relative to fset

// 	src := `package foo

// import (
// 	"fmt"
// 	"time"
// )

// func bar() string {
// 	root := []int{1, 1}
// 	// search from []int{1, 1} {
// 	// 	reject
// 	// }
// }`

// 	// Parse src but stop after processing the imports.
// 	f, err := parser.ParseFile(fset, "", src, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// // Print the imports from the file's AST.
// 	// for _, s := range f.Imports {
// 	// 	fmt.Println(s.Path.Value)
// 	// }
// 	// for _, s := range f.Decls {
// 	// 	fmt.Println(s)
// 	// }

// 	// ast.Inspect(f, func(n ast.Node) bool {
// 	// 	// Find Return Statements
// 	// 	ret, ok := n.(*ast.ReturnStmt)
// 	// 	if ok {
// 	// 		fmt.Printf("return statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
// 	// 		printer.Fprint(os.Stdout, fset, ret)
// 	// 		return true
// 	// 	}
// 	// 	return true
// 	// })

// 	ast.Inspect(f, func(n ast.Node) bool {
// 		// Find Return Statements
// 		ret, ok := n.(*ast.AssignStmt)
// 		if ok {
// 			// fmt.Printf("return statement found on line %d:\n\t\n", fset.Position(ret.Pos()).Line)
// 			// printer.Fprint(os.Stdout, fset, ret)
// 			log.Println(ret.Rhs[0.])
// 			// printer.Fprint(os.Stdout, fset, ret)
// 			return true
// 		}
// 		return true
// 	})

// }

package main

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"sort"
	"strings"

	"go/ast"
)

func main() {
	// Parse a single source file.
	const input = `
package fib

func fib(x int) []int {
	return []int{1,1}
	// search from []int{1,1} {

	// }
    return []int{1,1}
}`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "fib.go", input, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Type-check the package.
	// We create an empty map for each kind of input
	// we're interested in, and Check populates them.
	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	var conf types.Config
	pkg, err := conf.Check("fib", fset, []*ast.File{f}, &info)
	if err != nil {
		log.Fatal(err)
	}

	// Print package-level variables in initialization order.
	fmt.Printf("InitOrder: %v\n\n", info.InitOrder)

	// For each named object, print the line and
	// column of its definition and each of its uses.
	fmt.Println("Defs and Uses of each named object:")
	usesByObj := make(map[types.Object][]string)
	for id, obj := range info.Uses {
		posn := fset.Position(id.Pos())
		lineCol := fmt.Sprintf("%d:%d", posn.Line, posn.Column)
		usesByObj[obj] = append(usesByObj[obj], lineCol)
	}
	var items []string
	for obj, uses := range usesByObj {
		sort.Strings(uses)
		item := fmt.Sprintf("%s:\n  defined at %s\n  used at %s",
			types.ObjectString(obj, types.RelativeTo(pkg)),
			fset.Position(obj.Pos()),
			strings.Join(uses, ", "))
		items = append(items, item)
	}
	sort.Strings(items) // sort by line:col, in effect
	fmt.Println(strings.Join(items, "\n"))
	fmt.Println()

	fmt.Println("Types and Values of each expression:")
	items = nil
	for expr, tv := range info.Types {
		var buf bytes.Buffer
		posn := fset.Position(expr.Pos())
		tvstr := tv.Type.String()
		if tv.Value != nil {
			tvstr += " = " + tv.Value.String()
		}
		// line:col | expr | mode : type = value
		// fmt.Fprintf(&buf, "%2d:%2d | %-19s | %-7s : %s",
		// 	posn.Line, posn.Column, exprString(fset, expr),
		// 	mode(tv), tvstr)
		log.Println(posn.Line, posn.Column, expr, tv, tvstr)
		items = append(items, buf.String())
	}
	sort.Strings(items)
	fmt.Println(strings.Join(items, "\n"))
}

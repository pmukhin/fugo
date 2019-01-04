package parse

import (
	"fgo/ast"
	"fgo/lex"
	"reflect"
	"testing"
)

type testCase struct {
	input  string
	output *ast.File
}

func TestParser_Parse(t *testing.T) {
	testCases := []testCase{
		{
			input: "sum :: int -> int -> int",
			output: &ast.File{
				Decls: []ast.Decl{
					&ast.TypeDecl{
						Name: "sum",
						Args: []*ast.Ident{
							{"int", 7},
							{"int", 14},
						},
						DeclStart:  0,
						DeclEnd:    23,
						ReturnType: &ast.Ident{Value: "int", NamePos: 21},
					},
				},
			},
		},
		{
			input: "sum x y = x + y",
			output: &ast.File{
				Decls: []ast.Decl{
					&ast.FuncDecl{
						Name: "sum",
						Args: []*ast.Ident{
							{"x", 4},
							{"y", 6},
						},
						DeclStart: 0,
						DeclEnd:   14,
						Expr: &ast.BinaryOpExpr{
							Left:  &ast.ValExpr{Name: &ast.Ident{Value: "x"}, StartPos: 10},
							Right: &ast.ValExpr{Name: &ast.Ident{Value: "y"}, StartPos: 14},
							Op:    &ast.Ident{Value: "+", NamePos: 12},
						},
					},
				},
			},
		},
		{
			input: "add x y = x `sum` y",
			output: &ast.File{
				Decls: []ast.Decl{
					&ast.FuncDecl{
						Name: "add",
						Args: []*ast.Ident{
							{"x", 4},
							{"y", 6},
						},
						DeclStart: 0,
						DeclEnd:   14,
						Expr: &ast.BinaryOpExpr{
							Left:  &ast.ValExpr{Name: &ast.Ident{Value: "x"}, StartPos: 10},
							Right: &ast.ValExpr{Name: &ast.Ident{Value: "y"}, StartPos: 14},
							Op:    &ast.Ident{Value: "sum", NamePos: 12},
						},
					},
				},
			},
		},
	}

	for i, tt := range testCases {
		p := New(lex.New(tt.input))
		file, err := p.Parse()

		if err != nil {
			t.Errorf("error in #%d: %s", i, err.Error())
		}

		if !reflect.DeepEqual(file, tt.output) {
			t.Errorf("%v != %v in #%d", file, tt.output, i)
		}
	}
}

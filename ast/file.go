package ast

type File struct {
	//Doc        *CommentGroup   // associated documentation; or nil
	//Package    token.Pos       // position of "package" keyword
	//Name       *Ident          // package name
	Decls []Decl // top-level declarations; or nil
	//Scope      *Scope          // package scope (this file only)
	//Imports    []*ImportSpec   // imports in this file
	//Unresolved []*Ident        // unresolved identifiers in this file
	//Comments   []*CommentGroup // list of all comments in the source file
}

func (f File) String() string {
	decls := "["
	for i, d := range decls {
		decls += string(d)
		if i == len(decls)-1 {
			break
		}
		decls += ", "
	}
	decls += "]"

	return "ast.File { Decls =" + decls + " }"
}

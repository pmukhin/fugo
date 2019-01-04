package ast

type Node interface {
	Pos() int
	End() int
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

type Decl interface {
	Node
	declNode()
}

type TypeDecl struct {
	Name       string
	Args       []*Ident
	ReturnType *Ident
	DeclStart  int
	DeclEnd    int
}

func (t TypeDecl) Pos() int {
	return t.DeclStart
}

func (t TypeDecl) End() int {
	return t.DeclEnd
}

func (t TypeDecl) declNode() {}

type FuncDecl struct {
	Name      string
	Args      []*Ident
	DeclStart int
	DeclEnd   int
	Expr      Expr
}

func (f FuncDecl) Pos() int {
	panic("implement me")
}

func (f FuncDecl) End() int {
	panic("implement me")
}

func (f FuncDecl) declNode() {
	panic("implement me")
}

type Ident struct {
	Value   string
	NamePos int
}

func (i Ident) String() string {
	return i.Value
}

func (i Ident) Pos() int {
	panic("implement me")
}

func (i Ident) End() int {
	panic("implement me")
}

func (i Ident) exprNode() {
	panic("implement me")
}

type BinaryOpExpr struct {
	Left     Expr
	Right    Expr
	Op       *Ident
	StartPos int
}

func (b BinaryOpExpr) Pos() int {
	panic("implement me")
}

func (b BinaryOpExpr) End() int {
	panic("implement me")
}

func (b BinaryOpExpr) exprNode() {}

type ValExpr struct {
	Name *Ident

	StartPos int
}

func (v ValExpr) Pos() int {
	panic("implement me")
}

func (v ValExpr) End() int {
	panic("implement me")
}

func (v ValExpr) exprNode() {
	panic("implement me")
}

type FuncCall struct {
	Name *Ident
	Args []Expr

	StartPos int
}

func (f FuncCall) Pos() int {
	panic("implement me")
}

func (f FuncCall) End() int {
	panic("implement me")
}

func (f FuncCall) exprNode() {
	panic("implement me")
}


package parse

import (
	"fgo/ast"
	"fgo/lex"
	"fgo/token"
	"fmt"
)

type Parser struct {
	scanner  *lex.Scanner
	offset   int
	curToken token.Token
}

func New(scanner *lex.Scanner) *Parser {
	p := new(Parser)
	p.scanner = scanner
	p.next()

	return p
}

func (p *Parser) next() {
	p.curToken = p.scanner.Next()
}

func (p *Parser) Parse() (*ast.File, error) {
	var decls []ast.Decl

Parse:
	for {
		switch p.curToken.Typ {
		case token.Illegal:
			return nil, fmt.Errorf("unexpected token Illegal")
		case token.Eof:
			break Parse
		case token.Ident:
			decl, err := p.parseDecl()
			if err != nil {
				return nil, err
			}
			decls = append(decls, decl)
		}
	}

	f := &ast.File{
		Decls: decls,
	}

	return f, nil
}

func (p *Parser) parseDecl() (ast.Decl, error) {
	tokStart := p.curToken
	p.next() // eat name

	if p.curToken.Typ == token.Mean {
		return p.parseTypeDecl(tokStart)
	} else {
		return p.parseFuncDecl(tokStart)
	}
}

func (p *Parser) assertTokenIs(t token.Type) error {
	if p.curToken.Typ == t {
		return nil
	}
	return fmt.Errorf("expected token %d, got %v", t, p.curToken)
}

func (p *Parser) parseTypeDecl(tok token.Token) (ast.Decl, error) {
	p.next() // eat '::'

	var argDecls []*ast.Ident
	for {
		i, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		argDecls = append(argDecls, i)
		if p.curToken.Typ == token.RightArrow {
			p.next() // eat
		} else {
			break
		}
	}

	ret := argDecls[len(argDecls)-1]
	typeDecl := &ast.TypeDecl{
		Name:       tok.Str,
		Args:       argDecls[0 : len(argDecls)-1],
		ReturnType: ret,
		DeclStart:  tok.Pos,
		DeclEnd:    p.curToken.Pos - 1,
	}

	return typeDecl, nil
}

func (p *Parser) parseIdent() (*ast.Ident, error) {
	tok := p.curToken
	defer p.next() // eat token

	return &ast.Ident{Value: tok.Str, NamePos: tok.Pos}, nil
}

func (p *Parser) parseFuncDecl(tok token.Token) (*ast.FuncDecl, error) {
	args := p.parseArgs()
	err := p.assertTokenIs(token.Assign)

	if err != nil {
		return nil, err
	}

	p.next() // eat `=`
	expr, err := p.parseExpr()

	if err != nil {
		return nil, err
	}

	funcDecl := &ast.FuncDecl{
		Args:      args,
		DeclStart: tok.Pos,
		DeclEnd:   p.curToken.Pos - 1,
		Expr:      expr,
	}

	return funcDecl, nil
}

func (p *Parser) parseArgs() []*ast.Ident {
	var args []*ast.Ident

	for {
		if p.curToken.Typ == token.Ident {
			args = append(args, &ast.Ident{
				Value:   p.curToken.Str,
				NamePos: p.curToken.Pos,
			})
			p.next() // eat last Ident
		} else {
			break
		}
	}

	return args
}

/*
 * expr0: - funcCall x
 * 		  - funcCall . x + 1
 * expr1: x + 5
 * expr2: (+ 1)
 * expr3: 5
 * expr4: if x > 5 then 3 else 2
 */
func (p *Parser) parseExpr() (ast.Expr, error) {
	switch p.curToken.Typ {
	case token.Ident:
		return p.parseExpr0_Or_1()
	default:
		return nil, fmt.Errorf("unexpected token: %d", p.curToken.Typ)
	}
}

func (p *Parser) parseExpr0_Or_1() (ast.Expr, error) {
	initialIdent, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	switch p.curToken.Typ {
	case token.Operator:
		return p.parseBinOp(initialIdent)
	case token.Ident:
		return p.parseFuncCall(initialIdent)
	default:
		return nil, fmt.Errorf("unepxected token: %d", p.curToken.Typ)
	}
}

func (p *Parser) parseBinOp(ident *ast.Ident) (ast.Expr, error) {
	panic("not implemented")
}

func (p *Parser) parseFuncCall(ident *ast.Ident) (ast.Expr, error) {
	args, err := p.parseCallArgs()
	if err != nil {
		return nil, err
	}

	funcCall := &ast.FuncCall{
		Name:     ident,
		Args:     args,
		StartPos: ident.Pos(),
	}

	return funcCall, nil
}

func (p *Parser) parseCallArgs() ([]ast.Expr, error) {
	var args []ast.Expr

ParseAgain:
	switch p.curToken.Typ {
	case token.Ident:
		val, err := p.parseVal()
		if err != nil {
			return nil, err
		}
		args = append(args, val)
		goto ParseAgain
	//case token.OpenPar:
	//	return p.parseParenExpr()
	//case token.Dot:
	//	return p.parseDottedExpr()
	default:
		break ParseAgain
	}

	return args, nil
}

func (p *Parser) parseVal() (ast.Expr, error) {
	ident, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	return &ast.ValExpr{Name: ident, StartPos: ident.Pos()}, nil
}

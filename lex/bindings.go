package lex

import "fgo/token"

func tokenEof(pos int) token.Token {
	return token.Token{Typ: token.Eof, Pos: pos, Str: ""}
}

func tokenDot(pos int) token.Token {
	return token.Token{Typ: token.Dot, Pos: pos, Str: "."}
}

func tokenAssign(pos int) token.Token {
	return token.Token{Typ: token.Assign, Pos: pos, Str: "="}
}

func tokenIllegal(pos int, str string) token.Token {
	return token.Token{Typ: token.Illegal, Pos: pos, Str: str}
}

func tokenFloat(pos int, str string) token.Token {
	return token.Token{Typ: token.Float, Pos: pos, Str: str}
}

func tokenInteger(pos int, str string) token.Token {
	return token.Token{Typ: token.Integer, Pos: pos, Str: str}
}

func tokenIdent(pos int, str string) token.Token {
	return token.Token{Typ: token.Ident, Pos: pos, Str: str}
}

func tokenPar(pos int, open bool) token.Token {
	str := "("
	typ := token.OpenPar
	if !open {
		str = ")"
		typ = token.ClosePar
	}

	return token.Token{Typ: typ, Pos: pos, Str: str}
}

func tokenOp(pos int, lit string) token.Token {
	return token.Token{Typ: token.Operator, Pos: pos, Str: lit}
}

func tokenMean(pos int) token.Token {
	return token.Token{Typ: token.Mean, Pos: pos, Str: "::"}
}

func tokenRightArrow(pos int) token.Token {
	return token.Token{Typ: token.RightArrow, Pos: pos, Str: "->"}
}
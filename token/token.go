package token

type Type uint8

const (
	Illegal Type = iota
	Eof
	Ident
	Assign
	Dot
	Integer
	Float
	OpenPar
	ClosePar
	Operator
	Mean // ::
	RightArrow // ->
	OpenComment // {%
	CloseComment // %}
)

// Token
type Token struct {
	Typ Type
	Str string
	Pos int
}

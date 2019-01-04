package lex

import (
	"fgo/token"
	"reflect"
	"testing"
)

type testCase struct {
	input  string
	output []token.Token
}

func TestScanner_Next(t *testing.T) {
	tt := []testCase{
		{
			input: "main = putStrLn . 1 + 1",
			output: []token.Token{
				{
					token.Ident, "main", 0,
				},
				{
					token.Assign,
					"=",
					5,
				},
				{
					token.Ident,
					"putStrLn",
					7,
				},
				{
					token.Dot,
					".",
					16,
				},
				{
					token.Integer,
					"1",
					18,
				},
				{
					token.Operator,
					"+",
					20,
				},
				{
					token.Integer,
					"1",
					22,
				},
			},
		},
		{
			input: "x :: int -> string",
			output: []token.Token{
				{
					token.Ident, "x", 0,
				},
				{
					token.Mean, "::", 2,
				},
				{
					token.Ident, "int", 5,
				},
				{
					token.RightArrow, "->", 9,
				},
				{
					token.Ident, "string", 12,
				},
			},
		},
		{
			input: "(1.0)",
			output: []token.Token{
				{
					token.OpenPar, "(", 0,
				},
				{
					token.Float, "1.0", 1,
				},
				{
					token.ClosePar, ")", 4,
				},
			},
		},
		{
			input: "> >= < <= ==",
			output: []token.Token{
				{
					token.Operator, ">", 0,
				},
				{
					token.Operator, ">=", 2,
				},
				{
					token.Operator, "<", 5,
				},
				{
					token.Operator, "<=", 7,
				},
				{
					token.Operator, "==", 10,
				},
			},
		},
	}

	for i, tc := range tt {
		s := New(tc.input)
		tokens := make([]token.Token, 0)
		for {
			t := s.Next()
			if t.Typ == token.Eof || t.Typ == token.Illegal {
				break
			}
			tokens = append(tokens, t)
		}

		if reflect.DeepEqual(tokens, tc.output) {
			continue
		}

		t.Errorf("%v != %v in #%d", tokens, tc.output, i)
	}
}

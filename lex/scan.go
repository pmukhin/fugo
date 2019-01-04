package lex

import (
	"fgo/token"
	"unicode"
)

type Scanner struct {
	cur   rune
	len   int
	input []rune
	pos   int
}

func New(i string) *Scanner {
	s := new(Scanner)
	s.input = []rune(i)
	s.len = len(i)
	s.pos = -1
	s.cur = -1

	return s
}

func (s *Scanner) Next() token.Token {
	s.nextChar()
	s.skipWhitespace()

	switch s.cur {
	case -1:
		return tokenEof(s.pos)
	case '(':
		return tokenPar(s.pos, true)
	case ')':
		return tokenPar(s.pos, false)
	case '.':
		return tokenDot(s.pos)
	case ':':
		if s.peek() == ':' {
			posBackup := s.pos
			s.nextChar() // eat ':'
			return tokenMean(posBackup)
		} else {
			goto exitWithIllegal
		}
	case '-':
		if s.peek() == '>' {
			posBackup := s.pos
			s.nextChar() // eat '-'
			return tokenRightArrow(posBackup)
		} else {
			return s.parseOp()
		}
	case '=':
		if !isOp(s.peek()) {
			return tokenAssign(s.pos)
		} else {
			return s.parseOp()
		}
	default:
		switch true {
		case unicode.IsDigit(s.cur):
			return s.parseDigit()
		case unicode.IsLetter(s.cur):
			return s.parseIdent()
		case isOp(s.cur):
			return s.parseOp()
		default:
			goto exitWithIllegal
		}
	}

exitWithIllegal:
	return tokenIllegal(s.pos, string(s.cur))
}

func (s Scanner) peek() rune {
	s.nextChar()
	r := s.cur
	s.back()

	return r
}

func (s *Scanner) skipWhitespace() {
	for s.cur == ' ' {
		s.nextChar()
	}
}

func (s *Scanner) nextChar() {
	s.pos++
	if s.pos >= s.len {
		s.cur = -1
	} else {
		s.cur = s.input[s.pos]
	}
}

func (s *Scanner) parseDigit() token.Token {
	posBackup := s.pos
	isFloat := false

	var acc []rune
	for {
		if unicode.IsDigit(s.cur) {
			acc = append(acc, s.cur)
			s.nextChar()
		} else if s.cur == '.' {
			if isFloat {
				return tokenIllegal(s.pos, string(acc))
			} else {
				isFloat = true
				acc = append(acc, s.cur)
				s.nextChar()
			}
		} else {
			break
		}
	}

	s.back()
	if isFloat {
		return tokenFloat(posBackup, string(acc))
	}

	return tokenInteger(posBackup, string(acc))
}

func (s *Scanner) back() {
	s.pos--
	s.cur = s.input[s.pos]
}

func (s *Scanner) parseIdent() token.Token {
	posBackup := s.pos

	var acc []rune
	for isIdent(s.cur) {
		acc = append(acc, s.cur)
		s.nextChar()
	}

	s.back()

	return tokenIdent(posBackup, string(acc))
}

func (s *Scanner) parseOp() token.Token {
	posBackup := s.pos
	var acc []rune

	for isOp(s.cur) {
		acc = append(acc, s.cur)
		s.nextChar()
	}

	s.back()

	return tokenOp(posBackup, string(acc))
}

func isIdent(r rune) bool {
	if unicode.IsLetter(r) {
		return true
	}
	return false
}

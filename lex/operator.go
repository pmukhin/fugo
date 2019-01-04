package lex

var ex = struct{}{}
var opsCharMap = map[byte]struct{}{
	'=': ex,
	'+': ex,
	'@': ex,
	'<': ex,
	'>': ex,
	'&': ex,
	'^': ex,
}

func isOp(r rune) bool {
	if r > 255 {
		return false
	}

	_, ok := opsCharMap[byte(r)]
	return ok
}

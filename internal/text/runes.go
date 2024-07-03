package text

import "unicode/utf8"

func getRune(p []byte) rune {
	r, _ := utf8.DecodeRune(p)
	return r
}

var spaceRune rune = getRune([]byte{' '})
var crRune rune = getRune([]byte{'\r'})
var nlRune rune = getRune([]byte{'\n'})
var tabRune rune = getRune([]byte{'\t'})
var backtickRune rune = getRune([]byte{'`'})

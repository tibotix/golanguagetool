package text_processor

type Output func(rs []rune) []rune

func OutputNothing(rs []rune) []rune {
	return []rune{}
}
func OutputEcho(rs []rune) []rune {
	return rs
}
func OutputRunes(r ...rune) Output {
	return func(_ []rune) []rune {
		return r
	}
}
func OutputRune(r rune) Output {
	return OutputRunes(r)
}

package text

type LineBeginnings struct {
	lineBeginnings []int
}

func NewLineBeginningsFromArray(lineBeginnings []int) LineBeginnings {
	return LineBeginnings{
		lineBeginnings: lineBeginnings,
	}
}

func (l *LineBeginnings) LookupLine(pos int) int {
	return 0
}

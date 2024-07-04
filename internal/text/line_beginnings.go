package text

type LineBeginnings struct {
	data []int
}

func NewLineBeginningsFromArray(lineBeginnings []int) LineBeginnings {
	return LineBeginnings{
		data: lineBeginnings,
	}
}

func (l *LineBeginnings) LookupLine(pos int) int {
	low := 0
	high := len(l.data) - 1
	for low <= high {
		mid := low + (high-low)/2
		if l.data[mid] == pos {
			return mid + 1
		}
		if pos > l.data[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

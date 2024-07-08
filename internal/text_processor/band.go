package text_processor

type Band[E any] struct {
	// Note that it is not guaranteed that readHead is withing data bounds
	readHead int
	// Note that it is not guaranteed that writeHead is withing data bounds
	writeHead int
	data      []E
}

func NewBandWithCapacity[E any](capacity int) Band[E] {
	return Band[E]{
		readHead:  0,
		writeHead: 0,
		data:      make([]E, capacity),
	}
}
func NewBandWithData[E any](data []E) Band[E] {
	return Band[E]{
		readHead:  0,
		writeHead: 0,
		data:      data,
	}
}

func (b *Band[E]) areHeadsTogether() bool {
	return b.readHead == b.writeHead
}
func (b *Band[E]) extendIfNecessary(length int) {
	if length > len(b.data) {
		if length <= cap(b.data) {
			b.data = b.data[:length]
		} else {
			temp := make([]E, length)
			copy(temp, b.data)
			b.data = temp
		}
	}
}

func (b *Band[E]) MoveReadHead(n int) {
	if b.readHead+n < 0 {
		n = b.readHead
	}
	b.readHead += n
}
func (b *Band[E]) MoveWriteHead(n int) {
	if b.writeHead+n < 0 {
		n = b.writeHead
	}
	b.writeHead += n
}

func (b *Band[E]) Overwrite(value E) {
	b.extendIfNecessary(b.writeHead + 1)
	b.data[b.writeHead] = value
}
func (b *Band[E]) Write(value E) {
	b.Overwrite(value)
	b.writeHead++
}
func (b *Band[E]) OverwriteN(values []E) {
	b.extendIfNecessary(b.writeHead + len(values))
	copy(b.data[b.writeHead:b.writeHead+len(values)], values)
}
func (b *Band[E]) WriteN(values []E) {
	b.OverwriteN(values)
	b.writeHead += len(values)
}
func (b *Band[E]) Peek() E {
	b.extendIfNecessary(b.readHead + 1)
	return b.data[b.readHead]
}
func (b *Band[E]) PeekAhead(n int) E {
	b.extendIfNecessary(b.readHead + n + 1)
	return b.data[b.readHead+n]
}
func (b *Band[E]) PeekAheadUptoN(lower int, n int) []E {
	if b.readHead+lower >= len(b.data) {
		return []E{}
	}
	return b.data[b.readHead+lower : min(b.readHead+lower+n, len(b.data))]
}
func (b *Band[E]) Read() E {
	value := b.Peek()
	b.readHead++
	return value
}
func (b *Band[E]) PeekN(n int) []E {
	b.extendIfNecessary(b.readHead + n)
	return b.data[b.readHead : b.readHead+n]
}
func (b *Band[E]) PeekUptoN(n int) []E {
	b.extendIfNecessary(b.readHead + 1)
	return b.data[b.readHead:min(b.readHead+n, len(b.data))]
}
func (b *Band[E]) ReadN(n int) []E {
	values := b.PeekN(n)
	b.readHead += n
	return values
}

func (b *Band[E]) AllWritten() []E {
	b.extendIfNecessary(b.writeHead)
	return b.data[:b.writeHead]
}

func (b *Band[E]) AllRead() []E {
	b.extendIfNecessary(b.readHead)
	return b.data[:b.readHead]

}

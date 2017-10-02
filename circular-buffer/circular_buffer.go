package circular

import (
	"errors"
)

const testVersion = 4

// Buffer holds the properties of a ring buffer
type Buffer struct {
	b      []byte
	head   int
	length int
}

// NewBuffer takes a capacity (in bytes) and returns
// a new instance of a ring buffer.
func NewBuffer(size int) *Buffer {
	buf := Buffer{
		b: make([]byte, size),
	}
	return &buf
}

// ReadByte reads the oldest byte from the ring buffer
// and returns an error if the buffer is empty.
func (buf *Buffer) ReadByte() (byte, error) {

	if buf.length == 0 {
		return 0, errors.New("buffer empty")
	}

	tail := buf.advanceTail()
	c := buf.b[tail]

	return c, nil
}

// WriteByte writes a byte onto the head of the ring buffer
// and returns an error if the buffer is full.
func (buf *Buffer) WriteByte(c byte) error {
	if len(buf.b) == buf.length {
		return errors.New("")
	}

	buf.b[buf.head] = c
	buf.advanceHead()
	return nil
}

// Overwrite writes a byte onto the head of the ring buffer
// but overwrites the oldest bytes if the buffer is full.
func (buf *Buffer) Overwrite(c byte) {
	buf.b[buf.head] = c
	buf.advanceHead()
}

// Reset overwrites all the data in the buffer with zeros
// and resets both the read and write positions back to
// their initial values.
func (buf *Buffer) Reset() {
	for p := 0; p < len(buf.b); p++ {
		buf.b[p] = 0
	}
	buf.head = 0
	buf.length = 0
}

// advanceHead is a helper function that advances the
// index that points to the head of the ring buffer and
// returns it to the caller.
func (buf *Buffer) advanceHead() int {
	buf.head = (buf.head + 1) % len(buf.b)
	buf.length++
	if buf.length > len(buf.b) {
		buf.length = len(buf.b)
	}
	return buf.head
}

// advanceTail is a helper function that reduces the
// offset to the tail of the ring buffer and returns
// an index that points to the tail of the buffer.
func (buf *Buffer) advanceTail() int {
	tail := (buf.head + (len(buf.b) - buf.length)) % len(buf.b)
	buf.length--
	if buf.length < 0 {
		buf.length = 0
	}
	return tail
}

package circular

import (
	"errors"
)

const testVersion = 4

// Buffer holds the properties of a ring buffer
type Buffer struct {
	b chan byte
}

// NewBuffer takes a capacity (in bytes) and returns
// a new instance of a ring buffer.
func NewBuffer(size int) *Buffer {
	buf := Buffer{
		b: make(chan byte, size),
	}
	return &buf
}

// ReadByte reads the oldest byte from the ring buffer
// and returns an error if the buffer is empty.
func (buf *Buffer) ReadByte() (byte, error) {
	if len(buf.b) == 0 {
		return 0, errors.New("buffer empty")
	}
	c := <-buf.b
	return c, nil
}

// WriteByte writes a byte onto the head of the ring buffer
// and returns an error if the buffer is full.
func (buf *Buffer) WriteByte(c byte) error {
	if len(buf.b) == cap(buf.b) {
		return errors.New("buffer full")
	}
	buf.b <- c
	return nil

}

// Overwrite writes a byte onto the head of the ring buffer
// but overwrites the oldest bytes if the buffer is full.
func (buf *Buffer) Overwrite(c byte) {
	if len(buf.b) == cap(buf.b) {
		<-buf.b
	}
	buf.b <- c
}

// Reset overwrites all the data in the buffer with zeros
// and resets both the read and write positions back to
// their initial values.
func (buf *Buffer) Reset() {
	size := cap(buf.b)
	close(buf.b)
	buf.b = make(chan byte, size)
}

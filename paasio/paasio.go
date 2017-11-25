// Package paasio provides a wrapper for network connections and files that
// can report IO statistics including: the total number of bytes read/written,
// and the total number of read/write operations.
package paasio

import (
	"io"
	"sync"
)

// Counter is a struct that keeps track of read/write operations. It implements
// the ReadCounter, WriteCounter and ReadWriteCounter interfaces.
type Counter struct {
	w  io.Writer
	wc int64
	wo int
	r  io.Reader
	rc int64
	ro int
	m  sync.RWMutex
}

// Read wraps the basic Read method. It increments the counter by the number of
// bytes read and increments the number of read operations by one.
func (c *Counter) Read(b []byte) (int, error) {
	br, err := c.r.Read(b)
	c.m.Lock()
	c.rc += int64(br)
	c.ro++
	c.m.Unlock()
	return br, err
}

// ReadCount returns the total number of bytes read and the total number of read
// operations since the counter was created.
func (c *Counter) ReadCount() (n int64, nops int) {
	return c.rc, c.ro
}

// Write wraps the basic Read method. It increments the counter by the number of
// bytes written and increments the number of write operations by one.
func (c *Counter) Write(b []byte) (int, error) {
	bw, err := c.w.Write(b)
	c.m.Lock()
	c.wc += int64(bw)
	c.wo++
	c.m.Unlock()
	return bw, err
}

// WriteCount returns the total number of bytes written and the total number of
// write operations since the counter was created.
func (c *Counter) WriteCount() (n int64, nops int) {
	return c.wc, c.wo
}

// NewWriteCounter takes an io.Writer and wraps it in a counter. It returns a
// pointer to a Counter.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &Counter{
		w: w,
	}
}

// NewReadCounter takes an io.Reader and wraps it in a counter. It returns a
// pointer to a Counter.
func NewReadCounter(r io.Reader) ReadCounter {
	return &Counter{
		r: r,
	}
}

// NewReadWriteCounter takes an io.ReadWriter and wraps it in a counter. It returns a
// pointer to a Counter.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &Counter{
		w: rw,
		r: rw,
	}
}

package react

import "log"

const testVersion = 5

type Sheet struct{}

func (s Sheet) CreateInput(v int) InputCell {
	ic := C{}
	ic.SetValue(v)
	return &ic
}

func (s Sheet) CreateCompute1(c Cell, f func(a int) int) ComputeCell {
	cc := CC{}
	cc.c1 = c
	cc.SetValue(f) // TODO: store the function rather than value
	return &cc
}

func (s Sheet) CreateCompute2(c1, c2 Cell, f func(a, b int) int) ComputeCell {
	cc := CC{}
	return &cc
}

// C represents a cell.
type C struct {
	v int
}

// Value returns the value in a cell.
func (c *C) Value() int {
	//log.Printf("%+v", c)
	return c.v
}

// SetValue takes an integer and stores
// it as the value of an InputCell.
func (c *C) SetValue(v int) {
	log.Println("setting value to:", v)
	c.v = v
}

// CC represents a computed cell.
type CC struct {
	c1 Cell
	v  func(a int) int
}

func (cc *CC) SetValue(f func(a int) int) int {
	cc.v = f
	c := cc.c1
	return cc.v(c.Value())
}

func (cc *CC) Value() int {
	c := cc.c1
	return cc.v(c.Value())
}

type Callback struct {
	f func(int)
}

func (cb Callback) Cancel() {

}

func (c *CC) AddCallback(f func(int)) Canceler {
	cb := Callback{
		f: f,
	}
	return cb
}

func New() Reactor {
	s := Sheet{}
	return s
}

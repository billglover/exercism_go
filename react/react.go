package react

const testVersion = 5

type Sheet struct{}

func (s Sheet) CreateInput(v int) InputCell {
	ic := C{}
	ic.SetValue(v)
	return &ic
}

func (s Sheet) CreateCompute1(c Cell, f func(a int) int) ComputeCell {
	cc := CC1{}
	cc.c1 = c
	cc.f = f
	return &cc
}

func (s Sheet) CreateCompute2(c1, c2 Cell, f func(a, b int) int) ComputeCell {
	cc := CC2{}
	cc.c1 = c1
	cc.c2 = c2
	cc.f = f
	return &cc
}

// C represents a cell.
type C struct {
	v int
}

// Value returns the value in a cell.
func (c *C) Value() int {
	return c.v
}

// SetValue takes an integer and stores
// it as the value of an InputCell.
func (c *C) SetValue(v int) {
	c.v = v
}

// CC1 represents a computed cell.
type CC1 struct {
	f  func(a int) int
	c1 Cell
}

func (cc *CC1) SetValue(f func(a int) int) int {
	cc.f = f
	c := cc.c1
	return cc.f(c.Value())
}

func (cc *CC1) Value() int {
	c := cc.c1
	return cc.f(c.Value())
}

// CC2 represents a computed cell.
type CC2 struct {
	f  func(a, b int) int
	c1 Cell
	c2 Cell
}

func (cc *CC2) SetValue(f func(a, b int) int) int {
	cc.f = f
	return cc.Value()
}

func (cc *CC2) Value() int {
	return cc.f(cc.c1.Value(), cc.c2.Value())
}

type Callback struct {
	f func(int)
}

func (cb Callback) Cancel() {

}

func (c *CC1) AddCallback(f func(int)) Canceler {
	cb := Callback{
		f: f,
	}
	return cb
}

func (c *CC2) AddCallback(f func(int)) Canceler {
	cb := Callback{
		f: f,
	}
	return cb
}

func New() Reactor {
	s := Sheet{}
	return s
}

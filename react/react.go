package react

import "context"

const testVersion = 5

// New returns a new instance of a Reactor
func New() Reactor {
	s := Sheet{}
	return s
}

// Sheet represents a collection of cells
type Sheet struct{}

// CreateInput returns an instance of an InputCell
func (s Sheet) CreateInput(v int) InputCell {
	ic := C{}
	ic.SetValue(v)
	return &ic
}

// CreateCompute1 returns an instance of a ComputeCell that contains a function
// of one Cell
func (s Sheet) CreateCompute1(c Cell, f func(a int) int) ComputeCell {

	ctx := context.Background()
	fn := func() int {
		return f(c.Value())
	}

	u := Unit{}
	u.Init(ctx, fn, c.Register())

	return &u
}

// CreateCompute2 returns an instance of a ComputeCell that contains a function
// of two Cells
func (s Sheet) CreateCompute2(c1, c2 Cell, f func(a, b int) int) ComputeCell {

	ctx := context.Background()
	fn := func() int {
		return f(c1.Value(), c2.Value())
	}

	u := Unit{}
	u.Init(ctx, fn, c1.Register(), c2.Register())

	return &u
}

package react

// Sheet represents a collection of cells
type Sheet struct{}

func (s Sheet) CreateInput(v int) InputCell {
	ic := C{}
	ic.SetValue(v)
	return &ic
}

func (s Sheet) CreateCompute1(c Cell, f func(a int) int) ComputeCell {

	// subscribe to updates from our linked cell
	in := make(chan int, 1)
	c.Subscribe(in)

	cc := CC{
		f: func() int {
			return f(c.Value())
		},
	}
	return &cc
}

func (s Sheet) CreateCompute2(c1, c2 Cell, f func(a, b int) int) ComputeCell {
	cc := CC{
		f: func() int {
			return f(c1.Value(), c2.Value())
		},
	}
	return &cc
}

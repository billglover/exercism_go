package react

// Sheet represents a collection of cells
type Sheet struct{}

func (s Sheet) CreateInput(v int) InputCell {
	ic := C{}
	ic.SetValue(v)
	return &ic
}

func (s Sheet) CreateCompute1(c Cell, f func(a int) int) ComputeCell {

	cc := CC{
		f: func() int {
			return f(c.Value())
		},
	}
	cc.inputs = append(cc.inputs, c)
	return &cc
}

func (s Sheet) CreateCompute2(c1, c2 Cell, f func(a, b int) int) ComputeCell {
	cc := CC{
		f: func() int {
			return f(c1.Value(), c2.Value())
		},
	}
	cc.inputs = append(cc.inputs, []Cell{c1, c2}...)
	return &cc
}

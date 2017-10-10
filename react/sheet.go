package react

import "context"

// Sheet represents a collection of cells
type Sheet struct{}

func (s Sheet) CreateInput(v int) InputCell {
	ic := C{}
	ic.SetValue(v)
	return &ic
}

func (s Sheet) CreateCompute1(c Cell, f func(a int) int) ComputeCell {

	ctx := context.Background()
	fn := func() int {
		return f(c.Value())
	}
	//n1 := make(chan interface{}, 1)

	u := Unit{}
	u.Init(ctx, fn, c.Register())

	// cc := CC{
	// 	f: func() int {
	// 		return f(c.Value())
	// 	},
	// }
	// cc.inputs = append(cc.inputs, c)
	return &u
}

func (s Sheet) CreateCompute2(c1, c2 Cell, f func(a, b int) int) ComputeCell {
	// u := Unit{
	// 	fn: func() int {
	// 		return f(c1.Value(), c2.Value())
	// 	},
	// }
	// cc.inputs = append(cc.inputs, []Cell{c1, c2}...)
	// return &cc

	ctx := context.Background()
	fn := func() int {
		return f(c1.Value(), c2.Value())
	}


	u := Unit{}
	u.Init(ctx, fn, c1.Register(), c2.Register())

	return &u
}

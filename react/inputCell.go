package react

import "fmt"

// C represents a cell.
type C struct {
	v         int
	callbacks []Callback
}

// Value returns the value in a cell.
func (c *C) Value() int {
	return c.v
}

// SetValue takes an integer and stores
// it as the value of an InputCell.
func (c *C) SetValue(v int) {
	fmt.Println("setting value:", v)

	if c.Value() != v {
		c.v = v
		for _, cb := range c.callbacks {
			cc := cb.cc
			fmt.Println(cc)
			if cc.Value() != cb.oldValue {
				fmt.Println("calling callback:", cc.Value())
				cb.f(cc.Value())
				cb.oldValue = cc.Value()
			}
		}
	}
	c.v = v
}

// Subscribe registers a channel for subscription updates
func (c *C) Subscribe(cb Callback) {
	fmt.Println("storing callback in input cell")
	c.callbacks = append(c.callbacks, cb)
}

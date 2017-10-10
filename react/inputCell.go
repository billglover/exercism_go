package react

import "fmt"

// C represents a cell.
type C struct {
	v         int
	notifyChs []chan interface{}
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
		for _, cb := range c.notifyChs {
			fmt.Println("notifying:", true)
			cb <- true
		}
	}
	c.v = v
}

func (c *C) Register() chan interface{} {
	ch := make(chan interface{})
	c.notifyChs = append(c.notifyChs, ch)
	return ch
}

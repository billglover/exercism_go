package react

import "fmt"

// C represents a cell.
type C struct {
	v         int
	notifyChs []chan chan interface{}
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
			done := make(chan interface{}, 1)
			cb <- done

			// wait for subscibers to confirm
			// before returning
			<-done
		}
	}
	c.v = v
}

func (c *C) Register() chan chan interface{} {
	ch := make(chan chan interface{})
	c.notifyChs = append(c.notifyChs, ch)
	return ch
}

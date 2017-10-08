package react

import "fmt"

// C represents a cell.
type C struct {
	v    int
	subs []chan int
}

// Value returns the value in a cell.
func (c *C) Value() int {
	return c.v
}

// SetValue takes an integer and stores
// it as the value of an InputCell.
func (c *C) SetValue(v int) {
	c.v = v

	for i, ch := range c.subs {

		// if our subscriber isn't full, send a notification
		if len(ch) < cap(ch) {
			fmt.Println("notifying subscriber:", i)
			ch <- c.v
		}
	}
}

// Subscribe registers a channel for subscription updates
func (c *C) Subscribe(ch chan int) {
	c.subs = append(c.subs, ch)
}

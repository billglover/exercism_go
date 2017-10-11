package react

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

	if c.Value() != v {
		c.v = v
		for _, cb := range c.notifyChs {
			done := make(chan interface{}, 1)
			cb <- done
			<-done
		}
	}
	c.v = v
}

// Register allows dependent cells to register for
// update notifications when this cell's value changes
func (c *C) Register() chan chan interface{} {
	ch := make(chan chan interface{})
	c.notifyChs = append(c.notifyChs, ch)
	return ch
}

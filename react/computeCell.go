package react

// CC represents a computed cell.
type CC struct {
	f    func() int
	subs []chan int
}

// Value returns the value of a computed cell.
func (cc *CC) Value() int {
	return cc.f()
}

// Subscribe registers a channel for subscription updates
func (cc *CC) Subscribe(ch chan int) {
	cc.subs = append(cc.subs, ch)
}

type Callback struct{}

func (cb Callback) Cancel() {}

func (c *CC) AddCallback(f func(int)) Canceler {
	cb := Callback{}
	return cb
}

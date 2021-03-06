package react

import (
	"context"
	"sync"
)

// Unit represents a single compute unit, holding a
// function and one or more callbacks.
type Unit struct {
	fn          func() int   // stores the computed function
	subscribers map[*CB]bool // stores pointers to callback functions
	val         int          // stores the current value
	notifyChs   []chan chan interface{}
}

// Init initialises a unit of computation by providing it
// a function to compute and one or more notification
// channels on which it can receive change notifications.
// It also takes a context to allow cancelation.
func (u *Unit) Init(ctx context.Context, f func() int, updates ...<-chan chan interface{}) {

	// capture the initial value as we'll need it to determine changes
	u.fn = f

	// store the compute function that computes cell value
	u.val = u.fn()

	// create an empty map of subscribers
	u.subscribers = make(map[*CB]bool, 0)

	// create a single channel on which to receive update notifications
	notify := merge(updates...)

	// start a goroutine to wait for update notifications and trigger
	// callbacks when they arive
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case done := <-notify:
				if u.val != u.fn() {
					u.doCallbacks()
					u.Notify()
				}
				done <- true
			}
		}
	}()
}

// doCallbacks executes all callback functions currently
// registered with this compute unit.
func (u *Unit) doCallbacks() {
	for cb := range u.subscribers {
		cb.fn(u.Value())
	}
}

// Value returns the value of the compute unit by executing
// its stored function.
func (u *Unit) Value() int {
	u.val = u.fn()
	return u.val
}

// Notify indicates to any dependent Units that the value
// of this Unit has changed
func (u *Unit) Notify() {
	for _, cb := range u.notifyChs {
		done := make(chan interface{}, 1)
		cb <- done
		<-done
	}
}

// The merge function converts a list of channels to a
// single channel by starting a goroutine for each inbound
// channel that copies the values to the sole outbound
// channel. Once all the output goroutines have been
// started, merge starts one more goroutine to close the
// outbound channel after all sends on that channel are
// done.
// Source: https://blog.golang.org/pipelines
func merge(cs ...<-chan chan interface{}) <-chan chan interface{} {
	var wg sync.WaitGroup
	out := make(chan chan interface{})

	// Start an output goroutine for each input channel in
	// cs.  output copies values from c to out until c is
	// closed, then calls wg.Done.
	output := func(c <-chan chan interface{}) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output
	// goroutines are done.  This must start after the
	// wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// CB represents a callback function
type CB struct {
	fn     func(int)
	cancel func()
}

// Cancel removes a callback function
func (cb CB) Cancel() {
	cb.cancel()
}

// AddCallback is a stub function and needs to be updated
func (u *Unit) AddCallback(f func(int)) Canceler {
	cb := CB{}
	cb.fn = f
	cb.cancel = func() {
		delete(u.subscribers, &cb)
	}
	u.subscribers[&cb] = true
	return cb
}

// Register allows dependent cells to register for
// update notifications when this cell's value changes
func (u *Unit) Register() chan chan interface{} {
	ch := make(chan chan interface{})
	u.notifyChs = append(u.notifyChs, ch)
	return ch
}

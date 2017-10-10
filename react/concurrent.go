package react

import (
	"context"
	"fmt"
	"sync"
)

// Unit represents a single compute unit, holding a
// function and one or more callbacks.
type Unit struct {
	fn        func() int  // stores the computed function
	subs      []func(int) // stores a list of callbacks
	val       int         // stores the current value
	notifyChs []chan interface{}
}

// Init initialises a unit of computation by providing it
// a function to compute and one or more notification
// channels on which it can receive change notifications.
// It also takes a context to allow cancelation.
func (u *Unit) Init(ctx context.Context, f func() int, updates ...<-chan interface{}) {
	fmt.Println("starting compute unit")
	u.fn = f

	fmt.Println("updates to be monitored:", len(updates))

	// create a single channel on which to receive update
	// notifications
	notify := merge(updates...)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("terminating compute unit")
				return
			case <-notify:
				fmt.Println("notification recieved")
				u.doCallbacks()
			}
		}
	}()
}

// doCallbacks executes all callback functions currently
// registered with this compute unit.
func (u *Unit) doCallbacks() {
	fmt.Println("current callbacks:", len(u.subs))
	for i, cb := range u.subs {
		fmt.Println("callback:", i, u.Value())
		cb(u.Value())
		fmt.Println("callback:", i, "called")
	}
}

// Value returns the value of the compute unit by executing
// its stored function.
func (u *Unit) Value() int {
	if u.val != u.fn() {
		//fmt.Println("send notification that value has changed")
		u.val = u.fn()
	}
	return u.val
}

// The merge function converts a list of channels to a
// single channel by starting a goroutine for each inbound
// channel that copies the values to the sole outbound
// channel. Once all the output goroutines have been
// started, merge starts one more goroutine to close the
// outbound channel after all sends on that channel are
// done.
// Source: https://blog.golang.org/pipelines
func merge(cs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	// Start an output goroutine for each input channel in
	// cs.  output copies values from c to out until c is
	// closed, then calls wg.Done.
	output := func(c <-chan interface{}) {
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

// AddCallback is a stub function and needs to be updated
func (u *Unit) AddCallback(f func(int)) Canceler {
	fmt.Println("current callbacks:", len(u.subs))
	u.subs = append(u.subs, f)
	fmt.Println("current callbacks:", len(u.subs))
	return nil
}

func (u *Unit) Register() chan interface{} {
	ch := make(chan interface{})
	u.notifyChs = append(u.notifyChs, ch)
	return ch
}

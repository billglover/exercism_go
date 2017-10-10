package react

// import "fmt"

// // CC represents a computed cell.
// type CC struct {
// 	f      func() int
// 	inputs []Cell
// }

// // Value returns the value of a computed cell.
// func (cc *CC) Value() int {
// 	return cc.f()
// }

// // Subscribe registers a channel for subscription updates
// func (cc *CC) Subscribe(cb Callback) {
// }

// type Callback struct {
// 	oldValue int
// 	f        func(int)
// 	cc       ComputeCell
// }

// func (cb Callback) Cancel() {
// 	fmt.Println("callback cancelled")
// }

// func (cc *CC) AddCallback(f func(int)) Canceler {

// 	// how many linked cells do we need to add a callback to
// 	fmt.Println("linked cells:", len(cc.inputs))

// 	// create the callback
// 	cb := Callback{
// 		f:        f,
// 		oldValue: cc.Value(),
// 		cc:       cc,
// 	}

// 	// add our callback function to each of them
// 	for _, c := range cc.inputs {
// 		fmt.Println("trying to register callback")
// 		c.Subscribe(cb)
// 	}

// 	return cb
// }

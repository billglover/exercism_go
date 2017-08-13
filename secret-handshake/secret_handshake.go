package secret

const testVersion = 2

var actions = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake takes a decimal number and converts it to the
// appropriate sequence of events for a secret handshake.
func Handshake(n uint) (h []string) {

	index := 0

	// loop until we have processed all bits
	for n != 0 {

		// Increment the index until we have a 1 in the least
		// significant bit postion. The index can now be used
		// to look-up our secret handshake action.
		if n&1 == 1 {

			// handle the special case when we need to reverse
			// the list of actions
			if index == len(actions) {
				for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
					h[i], h[j] = h[j], h[i]
				}
				break
			}

			// append the action to our handshake
			h = append(h, actions[index])

		}
		index++

		// bit shift our integer to the right (divide by 2)
		n = n >> 1
	}

	return h
}

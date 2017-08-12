package gigasecond

import (
	"time"
)

const (
	testVersion = 4
	gigaSecond  = 1000000000 * time.Second
)

// AddGigasecond adds 1 gigasecond (1 * 10^9) to a time
// value and returns the result
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(gigaSecond)
	return t
}

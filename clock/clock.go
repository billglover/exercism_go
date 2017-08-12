package clock

import (
	"fmt"
)

const testVersion = 4

// Clock provides a structure for handling time based on a
// 24 hour clock. It stores time as a combination of hours
// and minutes.
type Clock struct {
	m int
	h int
}

// New takes in hours and minutes and returns a value of
// type Clock.
func New(hour, minute int) Clock {

	o, m := normMins(minute)
	h := normHours(hour + o)

	nc := Clock{
		m: m,
		h: h,
	}
	return nc
}

// String returns an hh:mm representation of the Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add increments the current Clock value by the specified
// number of minutes. It handles negative increments by
// subtracting minutes off the clock.
func (c Clock) Add(minutes int) Clock {
	o, m := normMins(c.m + minutes)
	h := normHours(c.h + o)

	nc := Clock{
		m: m,
		h: h,
	}
	return nc
}

func normHours(h int) (nh int) {
	nh = h % 24
	if nh < 0 {
		nh += 24
	}
	return nh
}

func normMins(m int) (nh, nm int) {
	nh = m / 60

	nm = m % 60
	if nm < 0 {
		nh--
		nm += 60
	}

	return nh, nm
}

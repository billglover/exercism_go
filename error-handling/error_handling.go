package erratum

import (
	"errors"
)

const testVersion = 2

// Use opens a resource and handles various errors
// that may occur when opening resources. Transient errors are
// retried, all others are returned.
func Use(ro ResourceOpener, input string) (err error) {

	var r Resource

	for {
		r, err = ro()

		// we should retry on transient errors
		if _, ok := err.(TransientError); ok == true {
			continue
		}

		// all other errors should be returned
		if err != nil {
			return err
		}

		// if we don't have an error, break out of the
		// retry loop
		break
	}

	// if we've opened the resource, close it
	defer r.Close()

	// if something triggers a panic, handle it
	defer func(r Resource) {
		if pr := recover(); pr != nil {

			// we need to Defrob on FrobErrors, whatever frobbing may be
			if e, ok := pr.(FrobError); ok == true {
				r.Defrob(e.defrobTag)
				err = e
			}

			// overwite the return value
			err = errors.New("meh")
		}
	}(r)

	r.Frob(input)
	return err
}

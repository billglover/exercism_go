// Package flatten provides a function to flatten nested lists
package flatten

// Flatten takes a nested list and returns a single flattened
// list with all values except nil/null.
func Flatten(nl interface{}) []interface{} {

	fl := make([]interface{}, 0)

	// Check that we have been passed a list.
	if l, ok := nl.([]interface{}); ok {

		// Loop over items in the list and handle anything that
		// is an int or another list. We don't explicitly need
		// to filter nil values as only int and lists are handled.
		for _, i := range l {
			switch i.(type) {
			case int:
				fl = append(fl, i.(int))
			case []interface{}:
				fi := Flatten(i)
				fl = append(fl, fi...)
			}
		}
	}
	return fl
}

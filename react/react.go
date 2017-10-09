package react

import "fmt"

const testVersion = 5

func New() Reactor {
	s := Sheet{}
	fmt.Println("---")
	return s
}

package sieve

import (
	"context"
)

const testVersion = 1

// Sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a
// given number.
func Sieve(max int) (primes []int) {

	ctx, cancel := context.WithCancel(context.Background())

	// create a source of numbers
	ch := make(chan int)
	go source(ctx, ch)

	for {
		p := <-ch

		if p > max {
			cancel()
			return primes
		}

		primes = append(primes, p)

		// filter out all future multiples of the prime number
		chx := make(chan int)
		go filter(ctx, ch, chx, p)
		ch = chx
	}

}

func source(ctx context.Context, ch chan<- int) {
	i := 2
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- i
			i++
		}
	}
}

func filter(ctx context.Context, in <-chan int, out chan<- int, prime int) {
	for {
		select {
		case <-ctx.Done():
			return
		case i := <-in:
			if i%prime != 0 {
				out <- i
			}
		}
	}
}

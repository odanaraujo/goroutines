package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primes(num int, ch chan int) {
	init := 2
	for i := 0; i < num; i++ {
		for prime := init; ; prime++ {
			if isPrime(prime) {
				ch <- prime
				init = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(ch)
}

func main() {
	ch := make(chan int, 30)

	go primes(cap(ch), ch)

	for value := range ch {
		fmt.Printf("%d - ", value)
	}

	fmt.Println("\nthe end")
}

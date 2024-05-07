package main

import (
	"fmt"
	"time"
)

func mutiplication(value int, ch chan int) {
	for i := 1; i <= value; i++ {
		time.Sleep(time.Second * time.Duration(1+i))
		ch <- value * i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go mutiplication(3, ch)

	for calc := range ch {
		fmt.Printf("multiplied value is %v\n", calc)
	}
}

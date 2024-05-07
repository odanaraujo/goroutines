package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1 //operação bloqueante
	fmt.Println("only after it has been read")
}

func main() {
	ch := make(chan int) //channel without buffer

	go routine(ch)
	fmt.Println(<-ch) // operation block
	fmt.Println("was read")
	fmt.Println(<-ch) // operation block
	fmt.Printf("end")
}

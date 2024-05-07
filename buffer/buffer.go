package main

import (
	"fmt"
	"strings"
)

func count(ch chan string, done chan bool, value string) {
	defer close(ch)
	for i, val := range value {
		if i == cap(ch) {
			break
		}
		ch <- strings.TrimSpace(strings.ToLower(string(val)))
	}
	done <- true
}

func main() {
	ch := make(chan string, 3)
	done := make(chan bool)
	go count(ch, done, "sport recife")

	go func() {
		<-done
		close(done)
	}()

	for value := range ch {
		fmt.Println(value)
	}
}

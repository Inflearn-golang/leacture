package main

import (
	"fmt"
	"time"
)

func work(s string, done chan string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s, "working", i, "hours")
		time.Sleep(time.Second)
	}
	done <- s
}

func main() {
	done := make(chan string, 8)
	work("programmer", done)
	work("designer", done)
	work("producer", done)
	work("marketer", done)

	go work("programmer", done)
	go work("designer", done)
	go work("producer", done)
	go work("marketer", done)

	for wait := range done {
		fmt.Println(wait)
	}
}

package main

import (
	"fmt"
	"time"
)

func longTimeRequest(word string) <-chan string {
	r := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		r <- word
	}()

	return r
}

func reverseWord(a, b string) string {
	return b + " " + a
}

// SimpleConcurrency is simple function
func SimpleConcurrency() {
	fmt.Println("Loading ...")
	a, b := longTimeRequest("world"), longTimeRequest("hello")
	fmt.Println(reverseWord(<-a, <-b))
}

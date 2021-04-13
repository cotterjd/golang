package main

import (
	"fmt"
	"time"
)

func printHello() {
	for {
		fmt.Println("hello world")
	}
}

// will print hello world for 2 seconds. Then will continue execution
// contrast with blocking.js
func main() {
	go printHello()
	time.Sleep(2 * time.Second)
	fmt.Println("The end")
}

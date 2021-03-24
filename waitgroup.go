package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// without sync package and wg.Add, wg.Done, and wg.Wait, println will never show. It will fire off routine and finish before it prints
	var msg = "Hello"
	wg.Add(1) // default is 0. This number represents how  many subroutines the main func is waiting on
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // decrements by 1 so back to 0 and the program can finish
	}(msg)
	wg.Wait() // waiting on subroutines
}


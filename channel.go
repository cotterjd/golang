package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
// channel is to data between goroutines
func main() {
	ch := make(chan int)
	wg.Add(2) // add wait time (passing 0, for example, won't give time for program to run. Will have empty output)
	go func() { // adding to channel must be in goroutine
		i := <- ch // take data out of channel
		fmt.Println(i)
		wg.Done() // channel will lock and program will panic without ending channel
	}()
	go func() { // taking from channel must be in go routine
		ch <- 42 // put data into channel
		wg.Done() // channel will lock and program will panic without ending channel
	}()
	wg.Wait() // wait for goroutines
}

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <- ch // take data out of channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // put data into channel
		wg.Done()
	}()
	wg.Wait()
}


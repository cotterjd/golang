package main

import (
	"fmt"
)

func main() {
	var f func(int) = func(a int) {
	  fmt.Println("Hello Go!", a)
	}
	myfunc(f)
}

func myfunc(x func(int)) {
  x(1)
}

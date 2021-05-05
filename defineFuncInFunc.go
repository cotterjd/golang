package main

import "fmt"

// wrong  syntax error: unexpected foo, expecting (
// func main () {
// 	func foo () {
// 		fmt.Println("foo")
// 	}
// 	foo()
// }

func main () {
	foo := func() {
		fmt.Println("foo")
	}
	foo()
}


package main

import (
	"fmt"
)

func main() {
	// Do not have to explicitly state that we're implimenting the interface
	// You impliment the interface by defining the method(s) of the interface
	type Writer interface {
		Write([]byte) (int, error)
	}

	type ConsoleWriter struct {}

	func (cw ConsoleWriter) Write(data []byte) (int, error) {
		n, err := fmt.Println(string(data))
		return n, err
	}
}


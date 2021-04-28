package main

import (
	"fmt"
)

type S struct{
	Foo string
}

func main() {
	s := S{
		Foo: "initial value",
	}
	changeFoo(s)
	fmt.Printf("%+v\n", s)
}

func changeFoo (s S) {
	fmt.Println("changing value of foo...")
	s.Foo = "new value"
}

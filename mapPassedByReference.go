package main

import (
	"fmt"
)

func main() {
	m := map[string]string{}
	m["foo"] = "initial value"
	changeFoo(m)
	fmt.Printf("%+v\n", m)
}

func changeFoo (m map[string]string) {
	m["foo"] = "value changed in function"
}

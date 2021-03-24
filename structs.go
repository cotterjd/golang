package main

import (
	"fmt"
)

func main() {

	// Wrong
	// type struct1 struct{
	// 	name string
	// 	age int
	// }
	// foo1 := {
	// 	name: "John Pertwee", 
	// 	age: 11,
	// }
	// fmt.Println(foo1) // syntax error: unexpected {, expecting expression

	// Wrong
	// type struct1 struct{
	// 	name string
	// 	age int
	// }
	// var foo1 struct1 = {
	// 	name: "John Pertwee",
	// 	age: 11,
	// }
	// fmt.Println(foo1) // syntax error: unexpected {, expecting expression

	// Wrong
	// type struct1 struct{ name string }
	// var foo1 struct1 = { name: "John Pertwee" }
	// fmt.Println(foo1) // syntax error: unexpected {, expecting expression

	// Wrong
	// type struct1 struct{
	// 	name string
	// 	age int
	// }
	// foo1 := struct1
	// fmt.Println(foo1) // type struct1 is not an expression

	// no commas on type. commas required after fields on struct value
	// Correct
	type struct1 struct{
		name string
		age int
	}
	foo1 := struct1{
		name: "Jon",
		age: 20,
	}
	fmt.Println(foo1)
}


package main

import (
	"fmt"
)

func main() {
	// Think of struct types as TypeScript types. They do not require commas.
	// Think of Go as always having the ESLint rule to enforce trailing commas on multiline objects(structs). Will error otherwise.

	// inline struct
	// Correct
	struct1 := struct{ name string }{ name: "Jon" }
	fmt.Println("inline struct with one property defined on one line:", struct1) 

	// Wrong
	// struct2 := struct{
	// 	name string
	// }{
	// 	name: "Jon"
	// }
	// fmt.Println(struct2.name)  // will error

	// Wrong
	// struct3 := struct{name string}{
	// 	name: "John Pertwee"
	// }
	// fmt.Println(struct3.name) // will error

	// Correct
	struct4 := struct{name string; age int}{name: "John Pertwee", age: 11}
	fmt.Println("inline struct with two properties defined on one line. Need ';' between properties on type. Need ',' between properties on value:", struct4)	


	// Wrong
	// struct5 := struct{name string, age int}{name: "John Pertwee", age: 11}
	// fmt.Println(struct5.age) // will error	

	// Correct
	struct6 := struct{
		name string
		age int
	}{name: "John Pertwee", age: 11}
	fmt.Println("inline struct with type defined on two lines. No comma after fields on type:", struct6)

	// Wrong
	// struct7 := struct{
	// 	name string
	// 	age int
	// }{
	// 	name: "John Pertwee",
	// 	age: 11
	// }
	// fmt.Println(struct7.age) // will error

	// Wrong
	// struct8 := struct{
	// 	name string
	// 	age int
	// }{
	// 	name: "John Pertwee"
	// 	age: 11
	// }
	// fmt.Println(struct8.age) // will error

	// Correct
	struct9 := struct{
		name string
		age int
	}{
		name: "John Pertwee",
		age: 11,
	}
	fmt.Println("inline struct with type and value both defined on two lines. No comma after fields on type. Commas after fields on value:", struct9) // will error

}


package main

import (
	"fmt"
)

func main() {
  value()
  reference()
  arrayEx()
  sliceEx()
}

func value() {
  fmt.Println("COPYING AN INT BY VALUE")
  foo := 3
  bar := foo 
  fmt.Printf("foo is 3. bar is set to foo so they have the same value. foo: %v, bar: %v\n", foo, bar)
  foo = 6
  fmt.Printf("foo has changed to 6. Now foo and bar have different values. foo: %v, bar: %v\n", foo, bar)
  fmt.Println("")
}
func reference() {
  fmt.Println("COPYING AN INT BY REFERENCE")
  var foo int = 3
  var bar *int = &foo 
  fmt.Printf("foo is a value. bar points to foo (var bar *int = &foo). bar represents the memory location of foo. foo: %v, bar: %v):\n", foo, bar)
  fmt.Printf("Adding the 'address of' operator (&) to foo will show it has the same memory location as bar. &foo: %v, bar: %v:\n", &foo, bar)
  fmt.Printf("Adding the dereferencing operator to bar will pull the value from foo. foo: %v, *bar: %v\n", foo, *bar)
  foo = 6
  fmt.Printf("Changing foo value to 6, also changes bar. foo: %v, *bar: %v\n", foo, *bar)
  fmt.Println("")
}
func arrayEx() {
  fmt.Println("COPYING AN ARRAY")
  arr := [3]int{1, 2, 3} // array
  arrCopy := arr
  fmt.Printf("arrCopy is set to arr(an array) so they are the same value. arr: %v, arrCopy: %v\n", arr, arrCopy)
  arr[1] = 42
  fmt.Printf("Element at index 1 changed to 42 for arr. arrCopy remains the same. arr: %v, arrCopy: %v\n", arr, arrCopy)
  fmt.Println("")
}
func sliceEx() {
  fmt.Println("COPYING A SLICE (same with maps")
  slic := []int{1, 2, 3} // slice
  slicCopy := slic 
  fmt.Printf("slicCopy is set to slic(a slice) so they are the same value. slic: %v, slicCopy: %v\n", slic, slicCopy)
  slic[1] = 42
  fmt.Printf("Element at index 1 changed to 42 for slic. sliceCopy changes as well since, as a slice, they both reference the same underlying array. slic: %v, slicCopy: %v\n", slic, slicCopy)
  fmt.Println("")
}

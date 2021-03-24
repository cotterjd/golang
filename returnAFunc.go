package main

import "fmt"

func main() {
    sqrtFunc:= getSqrtFunc()
    fmt.Println(sqrtFunc(3))
}

func getSqrtFunc() func(int) int {
    return func(x int) int {
        return x*x
    }
}

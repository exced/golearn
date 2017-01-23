package main

import "fmt"

const (
	pi     = 3.14
	myName = "Thomas"
)

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	var inString string
	fmt.Println("Enter your name")
	fmt.Scan(&inString)
	fmt.Println("Hello ", inString)
}

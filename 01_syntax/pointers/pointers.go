package main

import (
	"fmt"
)

// copy
func inc(a int) {
	a++
	fmt.Printf("%p\n", &a) // address in func inc
	fmt.Println(&a)        // address in func inc
}

// refer
func incp(a *int) {
	*a++
	fmt.Printf("%p\n", &a) // address in func inc
	fmt.Println(&a)        // address in func inc
}

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x... @a

	var b = &a
	fmt.Println(b)  // 0x... @a
	fmt.Println(*b) // 43

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42

	x := 1

	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main

	inc(x)
	fmt.Println("x ", x)

	incp(&x)
	fmt.Println("x p ", x)
}

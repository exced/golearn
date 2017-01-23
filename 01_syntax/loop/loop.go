package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println(i)
		if i >= 15 {
			break
		}
		i++
	}
}

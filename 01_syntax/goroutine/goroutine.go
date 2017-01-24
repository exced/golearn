package main

import (
	"fmt"
)

func main() {

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("countTo : ", i)
		}
	}()

	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("countFrom : ", i)
		}
	}()

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}

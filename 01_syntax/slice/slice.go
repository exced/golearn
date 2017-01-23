package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)

	// 0 is length - number of elements referred to by the slice
	// 3 is capacity - number of elements in the underlying array
	s2 := make([]int, 0, 3)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// dynamic arrays are arrays that double when fulled
	for i := 0; i < 50; i++ {
		s2 = append(s2, i)
		fmt.Println("Len:", len(s2), "Capacity:", cap(s2), "Value: ", s2[i])
	}

	s3 := []string{"Monday", "Tuesday"}
	s4 := []string{"Wednesday", "Thursday", "Friday"}

	s3 = append(s3, s4...)
	fmt.Println(s3)

	s3 = append(s3[:2], s4[3:]...)
	fmt.Println(s3)
}

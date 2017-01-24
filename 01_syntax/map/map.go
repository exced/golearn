package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]string)
	map1["hello"] = "thomas"
	map1["test"] = "ok"
	map1["hello"] = "test"
	fmt.Println(map1)

	map2 := map[int]string{}
	map2[0] = "hello"
	map2[1] = "world"
	fmt.Println(map2)

	map3 := map[int]string{
		0: "hello",
		1: "world",
	}
	delete(map3, 0)
	fmt.Println(map3)

	if val, exists := map2[0]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	if val, exists := map3[0]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
}

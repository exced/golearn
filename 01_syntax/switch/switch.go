package main

import (
	"fmt"
)

type nickname struct {
	name string
}

func switchOnType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case nickname:
		fmt.Println("nickname")
	}
}

func main() {
	switch "hello" {
	case "iam":
		fmt.Println("iam")
	case "hello", "your": // hello or your
		fmt.Println("hello")
		fallthrough
	case "further":
		fmt.Println("further")
	case "fallthrough":
		fmt.Println("fallthrough")
	default:
		fmt.Println("default")
	}

	/*
	   no default fallthrough
	   fallthrough is optional
	   -- you can specify fallthrough by explicitly stating it
	   -- break isn't needed like in other languages
	*/

	switchOnType(4)
	switchOnType("string")
	switchOnType(nickname{"thomas"})

}

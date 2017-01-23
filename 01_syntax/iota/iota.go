package main

import (
	"fmt"
)

// Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers. Because it can be used in expressions, it provides a generality beyond that of simple enumerations.

// The values of iota is reset to 0 whenever the reserved word const appears in the source (i.e. each const block) and increments by one after each ConstSpec e.g. each Line. This can be combined with the constant shorthand (leaving out everything after the constant name) to very concisely define related constants.
const (
	a = iota // 0
	b        // 1
	c        // 2
)

const (
	d = iota // 0
	e = iota // 1
	f = iota // 2
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}

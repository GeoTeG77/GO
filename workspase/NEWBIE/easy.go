package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := adder(a, b)
	fmt.Println(c)
	d := subtractor(a, b)
	fmt.Println(d)
	e := multiplier(a, b)
	fmt.Println(e)
	f := divider(a, b)
	fmt.Println(f, c, d, e)
}

func adder(a, b int) int {
	return a + b
}

func subtractor(a, b int) int {
	return a - b
}

func multiplier(a, b int) int {
	return a * b
}

func divider(a, b int) int {
	return a / b
}

package main

import "fmt"

func math_func(a int, b int) (sum int, sub int, mul int, div int) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return
}

func main() {
	s, sb, m, d := math_func(3, 5)
	fmt.Println(s, sb, m, d)

}

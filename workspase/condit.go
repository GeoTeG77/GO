package main

import "fmt"

func main() {
	num := 1
	if num > 2 {
		num += 3
		fmt.Print(num)
	} else {
		num -= 2
		fmt.Print(num)
	}
}

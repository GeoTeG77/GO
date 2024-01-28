package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, disk, exone, extwo float64
	fmt.Println("Your equation is : ax^2+bx+c=0")
	fmt.Println("Enter a,b,c,")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	disk = (b * b) - (4 * a * c)
	if disk < 0 {
		fmt.Println("No solution")
	} else if disk > 0 {
		exone = ((-b) + math.Sqrt(disk)) / (2 * a)
		extwo = ((-b) - math.Sqrt(disk)) / (2 * a)
		fmt.Println("There are two solution of your equation X1 = ", exone, "and X2 = ", extwo)
	} else if disk == 0 {
		exone = (-b) / (2 * a)
		fmt.Println("There is one solution of your eqation X = ", exone)
	}
	fmt.Scan(&a)
}

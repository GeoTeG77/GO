package main

import (
	"fmt"
)

var countOfMultipleNumbers, multiple int

func main() {
	fmt.Println("Enter a count of MultipleNumbers")
	fmt.Scanf("%d\n", &countOfMultipleNumbers)
	fmt.Println("Enter a multiple")
	fmt.Scanf("%d\n", &multiple)
	countBy(countOfMultipleNumbers, multiple)
}

func countBy(n int, x int) {
	arrayOfMultipleNumbers := []int{}
	
	for i := x; i <= n; i++ {
		if i%x == 0 {
			arrayOfMultipleNumbers = append(arrayOfMultipleNumbers, i)
		}
		
		fmt.Println(arrayOfMultipleNumbers)
	}

}

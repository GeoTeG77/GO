package main

import (
	"fmt"
	"math"
)

func main() {
	var symbPos int
	var condition, findNumber, number, toArr, pow float64
	arrayOfChair := []float64{}
	fmt.Println("Enter number")
	fmt.Scanf("%d\n", &number)
	condition = number
	findNumber = 1
	for pow := 1; condition >= 0; pow++ {
		toArr = math.Mod(condition, 10)
		arrayOfChair = append(arrayOfChair, toArr)
		condition = condition / 10
	}
	fmt.Println("Enter a symbol position for check")
	fmt.Scanf("%d\n", &symbPos)
	if symbPos > int(pow) {
		fmt.Println("Error")
	} //else
	for i := 0; i == int(pow); i++ {
		findNumber *= math.Pow(arrayOfChair[i],pow)
		fmt.Println(findNumber)
	}
}

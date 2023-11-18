package main

import (
	"fmt"
	"math"
)

func main() {

	var sum float64 = 0
	markes := [5]float64{12, 15, 23, 44, 55}
	markes[4] = 12
	fmt.Println(markes[4])

	for i := 0; i < len(markes); i++ {
		sum += markes[i]
	}
	var result float64 = sum / float64(len(markes))
	fmt.Println(math.Round(result))
}

package main

import (
	"fmt"
)

func main() {
	var dogsCount int
	fmt.Println("How many dogs you have?")
	fmt.Scanf("%d\n", &dogsCount)
	dogs := [4]string{"Hardly any", "More than a handful!", "It is a lot of dogges", "Woah, 101 DOLMATINOS"}
	if dogsCount <= 10 {
		fmt.Println(dogs[0])
	}
	if dogsCount > 10 {
		fmt.Println(dogs[1])
	}
	if dogsCount >= 50 {
		fmt.Println(dogs[2])
	}
	if dogsCount == 101 {
		fmt.Println(dogs[3])
	}

}

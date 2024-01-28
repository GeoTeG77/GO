package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 2, 5, 7, 4}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, el := range slice {
		fmt.Println("%d\n", el)
		slice = append(slice, 0)
	}
	slice[0] = 11
	wslice := []string{"я", "в", "а", "р"}
	sort.Ints(slice)
	sort.Strings(wslice)
	fmt.Println(slice[1])
	fmt.Println(slice)
	fmt.Println(wslice)
}

package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollars": 1000,
		"euros":   2000,
		"apples":  3,
	}
	fmt.Println(money)
	fmt.Println(money["dollars"])

	normmap := map[string]int{
		"nuga":   200,
		"kuga":   300,
		"yuersh": 600,
	}
	normmap["yuersh"] = 700
	delete(normmap, "kuga")
	fmt.Println(normmap)
}

package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
	score    []int
}

func (u User) getHighScore() int {
	high := 0

	for _, sc := range u.score {
		if high < sc {
			high = sc
		}
	}
	return high
}

func main() {

	us1 := User{"John", 23, "pass", []int{23, 67, 87, 642, 21}}
	fmt.Println(us1.getHighScore())
}

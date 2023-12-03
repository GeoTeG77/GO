package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	ch2 := make(chan int)
	go say("Hello GO", ch, ch2)

	data1 := <-ch
	data2 := <-ch2
	close(ch)
	close(ch2)
	fmt.Println(data1, data2)
}
func say(greet string, ch chan string, ch2 chan int) {
	fmt.Println(greet)
	ch2 <- 45
	ch <- "HALLO"

}

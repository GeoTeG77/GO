package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello GO")
	time.Sleep(2 * time.Second)
}

func say(greet string) {
	fmt.Println(greet)

}

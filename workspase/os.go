package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")

	if err != nil {
		fmt.Println("Error - ", err)
		os.Exit(1)
	}

	defer file.Close()

	data := "Text to file. Hello\n From\n Go\n\tTab"

	file.WriteString(data)
	file_data, _ := os.ReadFile(file.Name())
	fmt.Println(string(file_data))
	os.Remove("text.txt")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	cont := os.Args
	file, _ := os.Open("quest8.txt")
	arr := make([]byte, 14)
	file.Read(arr)
	file.Close()
	if len(cont) > 2 {
		fmt.Println("Too many arguments")
	}
	if len(cont) == 1 {
		fmt.Println("File name missing")
	}
	if len(cont) == 2 {
		fmt.Println(string(arr))
	}
}

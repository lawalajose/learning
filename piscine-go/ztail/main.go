package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}
		n = n*10 + int(ch-'0')
	}
	return n
}

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		return
	}

	n := atoi(os.Args[2])
	if n < 0 {
		fmt.Println("Invalid byte count")
		return
	}

	files := os.Args[3:]
	exitCode := 0

	for i, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			exitCode = 1
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", file)
		}

		if len(data) > n {
			data = data[len(data)-n:]
		}
		fmt.Print(string(data))
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}

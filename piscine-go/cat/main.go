package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	larg := len(arg)
	if larg == 1 {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				PrintString(err.Error() + "\n")
				break
			}
			PrintString(string(buf[:n]))
		}
		return
	}
	for i := 1; i < larg; i++ {
		argg := os.Args[i]
		PrintFile(argg)
	}
}

func PrintString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		PrintString("ERROR: " + err.Error() + "\n")
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		PrintString(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			PrintString("ERROR: " + err.Error() + "\n")
			break
		}
	}
}

package main

import "os"

func main() {
	if len(os.Args) != 4 {
		return
	}

	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, ok1 := atoi(aStr)
	b, ok2 := atoi(bStr)
	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		if willOverflowAdd(a, b) {
			return
		}
		printNbr(a + b)
	case "-":
		if willOverflowSub(a, b) {
			return
		}
		printNbr(a - b)
	case "*":
		if willOverflowMul(a, b) {
			return
		}
		printNbr(a * b)
	case "/":
		if b == 0 {
			printStr("No division by 0")
			return
		}
		printNbr(a / b)
	case "%":
		if b == 0 {
			printStr("No modulo by 0")
			return
		}
		printNbr(a % b)
	default:
		return
	}
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte{'\n'})
}

func printNbr(n int) {
	if n == 0 {
		os.Stdout.Write([]byte{'0', '\n'})
		return
	}
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte(n%10 + '0')}, digits...)
		n /= 10
	}
	digits = append(digits, '\n')
	os.Stdout.Write(digits)
}

func atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}
	num := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		d := int(s[i] - '0')
		if willOverflowMul(num, 10) || willOverflowAdd(num*10, d) {
			return 0, false
		}
		num = num*10 + d
	}
	return num * sign, true
}

const (
	MaxInt = 9223372036854775807
	MinInt = -9223372036854775808
)

func willOverflowAdd(a, b int) bool {
	if (b > 0 && a > MaxInt-b) || (b < 0 && a < MinInt-b) {
		return true
	}
	return false
}

func willOverflowSub(a, b int) bool {
	if (b < 0 && a > MaxInt+b) || (b > 0 && a < MinInt+b) {
		return true
	}
	return false
}

func willOverflowMul(a, b int) bool {
	if a == 0 || b == 0 {
		return false
	}
	if a == -1 && b == MinInt {
		return true
	}
	if b == -1 && a == MinInt {
		return true
	}
	if a > 0 && b > 0 && a > MaxInt/b {
		return true
	}
	if a < 0 && b < 0 && -a > MaxInt/-b {
		return true
	}
	if a > 0 && b < 0 && b < MinInt/a {
		return true
	}
	if a < 0 && b > 0 && a < MinInt/b {
		return true
	}
	return false
}

package main

import "fmt"

func main() {

	input := 67
	fmt.Println(Atoi(input))

}

func Atoi(n int) string {
	var result string
	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}
	if n == 0 {
		return "0"
	}
	for n > 0 {
		digit := n % 10

		result = string('0'+rune(digit)) + result
		n = n / 10
	}

	if isNegative {
		return "-" + result
	}

	return result

}

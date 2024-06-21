package main

import "fmt"

func main() {
	input := "love"
	fmt.Println(WordsToMarks(input))

}

func WordsToMarks(s string) int {

	var result int
	for _, char := range s {

		result += int(char - 'a' + 1)
	}
	return result

}

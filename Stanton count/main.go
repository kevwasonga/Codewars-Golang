package main

import "fmt"

func main() {
	input := []int{1, 4, 3, 2, 2, 1, 2, 3, 2}

	fmt.Println(StantonMeasure(input))
}

func StantonMeasure(arr []int) int {
	var count int
	// find number of times n appears
	for _, n := range arr {
		if n == 1 {
			count++
		}
	}

	// find stanton
	var stanton int

	for _, m := range arr {
		if m == count {
			stanton++
		}
	}

	return stanton
}

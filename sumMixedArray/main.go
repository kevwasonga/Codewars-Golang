package main

import (
	"fmt"
	"strconv"
)

// sumArray calculates the sum of both integer and string elements in the array
func sumArray(arr []interface{}) int {
	sum := 0

	for _, value := range arr {
		switch v := value.(type) {
		case int: // If the element is an integer, just add it to the sum
			sum += v
		case string: // If the element is a string, convert it to an integer and then add it
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			sum += num
		default: // If the type is neither int nor string, just ignore it
			fmt.Println("Unsupported type found:", v)
		}
	}

	return sum
}

func main() {
	// Example array with both integers and strings
	arr := []interface{}{1, "2", 3, "4", 5, "10", "abc", 12}

	// Call the sumArray function to calculate the sum
	result := sumArray(arr)

	// Print the result
	fmt.Println("The sum is:", result)
}

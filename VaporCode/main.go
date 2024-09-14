package main

import "fmt"

func main() {
	input := "Lets go to the movies"

	fmt.Println(Vaporcode(input))
}

func Vaporcode(s string) string {
	var res1 string
	// convert to capital where necessary

	for _, char := range s {

		if char == ' ' {
			continue
		}
		if char >= 'a' && char <= 'z' {
			res1 += string(char - 32)
		} else {
			res1 += string(char)
		}
	}

	// return res1
	fmt.Println(res1)

	// add spaces
	var res2 string

	for _, val := range res1 {
		res2 += "  " + string(val)
	}

	//	return res2

	// remove trailing spaces

	start := 0
	end := len(res2) - 1

	for start <= end && res2[start] == ' ' {
		start++
	}

	for end >= start && res2[end] == ' ' {
		end--
	}

	return res2[start : end+1]
}

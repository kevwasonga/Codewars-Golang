//a go program that return a string with exactly three spaces between them with no leading or trailing spaces
package main

import "fmt"

func main() {
	word := "    hello                                     world, are you reaady!!!!!!                         "
	fmt.Println(expandStr(word))
}

func expandStr(s string) string {
	// Initialize start and end indices
	start := 0
	end := len(s) - 1

	// Trim leading spaces
	for start <= end && s[start] == ' ' {
		start++
	}

	// Trim trailing spaces
	for end >= start && s[end] == ' ' {
		end--
	}

	// Return the trimmed string
	s1 := s[start : end+1]
	var result string

	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' && s1[i+1] == ' ' {
			continue
		} else if s1[i] == ' ' && s1[i+1] != ' ' {
			result += "   " // include the exact number of spaces neeeded between the words, in this scenario we add three spaces
		} else {
			result += string(s1[i])
		}
	}
	return result
}

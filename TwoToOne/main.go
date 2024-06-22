package main

import "fmt"

func main() {
	s1 := "asalaam"
	s2 := "alaikum"

	fmt.Println(TwoToOne(s1, s2))
}

func TwoToOne(s1 string, s2 string) string {
	seen := make(map[rune]bool) // Map to track seen letters
	var combined []rune         // Slice to store combined characters

	// Append all characters from s2 to s1
	for _, char := range s2 {
		s1 += string(char)
	}

	// Iterate over each character in s1 (which now includes s2)
	for _, char := range s1 {
		if !seen[char] {
			combined = append(combined, char)
			seen[char] = true
		}
	}

	// Sort combined characters alphabetically
	for i := 0; i < len(combined)-1; i++ {
		for j := i + 1; j < len(combined); j++ {
			if combined[i] > combined[j] {
				combined[i], combined[j] = combined[j], combined[i]
			}
		}
	}

	return string(combined)
}

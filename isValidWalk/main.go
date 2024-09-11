package main

import "fmt"

func main() {
	input := []rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}
	fmt.Println(IsValidWalk(input))
}

func IsValidWalk(walk []rune) bool {
	var blocks [][]rune

	var count int
	for _, char := range walk {
		if char == 'n' {
			count++
			// consider n,s,w,e as coordinates on  a graph that move by one step at a time
			blocks = append(blocks, []rune{0, 1})
		} else if char == 's' {
			count++
			blocks = append(blocks, []rune{0, -1})
		} else if char == 'w' {
			count++
			blocks = append(blocks, []rune{-1, 0})
		} else if char == 'e' {
			count++
			blocks = append(blocks, []rune{1, 0})
		} else {
			continue
		}
	}

	// fmt.Println(count) , count should return the number of valid characters found in the slice of rune, which in this case are treated as minutes

	var finalBLock [][]rune

	currentBlock := []rune{0, 0}

	for _, move := range blocks {

		Xmove := currentBlock[0] + move[0]
		Ymove := currentBlock[1] + move[1]

		currentBlock = []rune{Xmove, Ymove}

		finalBLock = append(finalBLock, currentBlock)

	}
//printing final block would give us all the destinations visited
	fmt.Println(finalBLock)

	expected := []rune{0, 0}

	if currentBlock[0] == expected[0] && currentBlock[1] == expected[1] && count == 10 {
		return true
	}
	return false

	
}
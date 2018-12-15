package day5

import (
	"fmt"
)

func Star2() {
	startingPolymer := []byte(input)
	shortestPolymer := len(startingPolymer)

	for c := byte('a'); c <= byte('z'); c++ {
		polymer := make([]byte, len(startingPolymer))
		copy(polymer, startingPolymer)
		for i := 0; i < len(polymer); i++ {
			if polymer[i] == c || polymer[i] == c-32 {
				polymer = append(polymer[:i], polymer[i+1:]...)
				i -= 1
			}
		}
		for i := 0; i < len(polymer)-1; i++ {
			if shouldAnnihilate(polymer[i], polymer[i+1]) {
				if i == 0 {
					polymer = polymer[2:]
					i = -1
				} else {
					polymer = append(polymer[:i], polymer[i+2:]...)
					i -= 2
				}
			}
		}

		if len(polymer) < shortestPolymer {
			shortestPolymer = len(polymer)
		}
	}

	fmt.Println("day5 star2:", shortestPolymer)
}

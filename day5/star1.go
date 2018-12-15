package day5

import "fmt"

func shouldAnnihilate(v1, v2 byte) bool {
	return v1-v2 == 32 || v2-v1 == 32
}

func Star1() {
	polymer := []byte(input)
	// fmt.Println(string(polymer))

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

	fmt.Println("day5 star1:", len(polymer))
}

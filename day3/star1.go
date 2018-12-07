package day3

import (
	"fmt"
	"strings"
)

func Star1() {
	fabricClams := make(map[int]map[int]int)
	claims := strings.Split(input, "\n")
	for _, claim := range claims {
		claimId, x, y, w, h := 0, 0, 0, 0, 0
		fmt.Sscanf(claim, "#%d @ %d,%d: %dx%d\n", &claimId, &x, &y, &w, &h)
		for c := 0; c < w; c++ {
			_, ok := fabricClams[c+x]
			if !ok {
				fabricClams[c+x] = make(map[int]int)
			}
			for r := 0; r < h; r++ {
				fabricClams[c+x][r+y] += 1
			}
		}
	}

	sharedSquareInches := 0
	for _, v := range fabricClams {
		for _, vv := range v {
			if vv > 1 {
				sharedSquareInches += 1
			}
		}
	}

	fmt.Println("day3 star1:", sharedSquareInches)
}

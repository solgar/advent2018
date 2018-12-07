package day3

import (
	"fmt"
	"strings"
)

func Star2() {
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

ClaimCheck:
	for _, claim := range claims {
		claimId, x, y, w, h := 0, 0, 0, 0, 0
		fmt.Sscanf(claim, "#%d @ %d,%d: %dx%d\n", &claimId, &x, &y, &w, &h)
		for c := 0; c < w; c++ {
			for r := 0; r < h; r++ {
				if fabricClams[c+x][r+y] > 1 {
					continue ClaimCheck
				}
			}
		}
		fmt.Println("day3 star2:", claimId)
	}
}

package day2

import (
	"fmt"
	"strings"
)

func Star1() {
	ids := strings.Split(input, "\n")
	counts := make(map[int]int)
	for _, id := range ids {
		chars := make(map[string]int)
		for _, c := range id {
			chars[string(c)] += 1
		}
		found2 := false
		found3 := false
		for _, v := range chars {
			if !found2 && v == 2 {
				found2 = true
				counts[2] += 1
			} else if !found3 && v == 3 {
				found3 = true
				counts[3] += 1
			}
		}
	}
	fmt.Println("day2 star1:", counts[2]*counts[3])
}

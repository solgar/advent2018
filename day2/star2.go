package day2

import (
	"fmt"
	"strings"
)

func Star2() {
	ids := strings.Split(input, "\n")
	for k, v := range ids {
		for kk, vv := range ids {
			if k == kk {
				continue
			}
			differences := 0
			differentIdx := 0
			for idx := range v {
				if v[idx] != vv[idx] {
					differences += 1
				}
				if differences == 1 {
					differentIdx = idx - 1
				}
			}
			if differences == 1 {
				fmt.Println("day2 star2:", v[:differentIdx]+v[differentIdx+1:])
				return
			}
		}
	}
}

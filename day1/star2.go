package day1

import (
	"advent2018/util"
	"fmt"
	"strconv"
	"strings"
)

func Star2() {
	currFreq := 0
	ops := strings.Split(input, "\n")
	freqOccurence := make(map[int]int)
	for {
		for _, op := range ops {
			i, err := strconv.Atoi(op)
			util.PanicIfError(err)
			currFreq += i
			freqOccurence[currFreq] += 1
			if freqOccurence[currFreq] == 2 {
				fmt.Println("day1 task2:", currFreq)
				return
			}
		}
	}
}

package day1

import (
	"advent2018/util"
	"fmt"
	"strconv"
	"strings"
)

func Star1() {
	freq := 0
	ops := strings.Split(input, "\n")
	for _, op := range ops {
		i, err := strconv.Atoi(op)
		util.PanicIfError(err)
		freq += i
	}
	fmt.Println("day1 task1:", freq)
}

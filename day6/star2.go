package day6

import (
	"fmt"
	"strings"
)

const (
	SAFE_AREA_THRESHOLD = 10000
)

func Star2() {
	inputLines := strings.Split(input, "\n")
	points = make([]*point, len(inputLines))
	maxX, maxY := -1, -1

	// create starting points references and get max width and height of field
	for i, v := range inputLines {
		p := &point{0, 0}
		points[i] = p
		fmt.Sscanf(v, "%d, %d", &p.x, &p.y)
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// add 1 to actually contain bounduary points
	maxX += 1
	maxY += 1

	// create field data
	field = make([][]fieldInfo, maxX)
	for i := 0; i < maxX; i++ {
		field[i] = make([]fieldInfo, maxY)
	}

	// calculate distances for each cell in field and calculate safe area
	safeArea := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			dst := 0
			for _, p := range points {
				dst += manhattan(point{x, y}, *p)
			}
			if dst < SAFE_AREA_THRESHOLD {
				safeArea += 1
			}
		}
	}

	fmt.Println("day6 star2:", safeArea)
}

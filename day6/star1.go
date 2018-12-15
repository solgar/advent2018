package day6

import (
	"fmt"
	"strings"
)

const (
	MIN_DST_UNDEFINED      = int((^uint(0)) >> 1)
	MIN_DST_STARTING_POINT = -1
)

type point struct {
	x, y int
}

type fieldInfo struct {
	minDistance        int
	closestCoordinates []*point
}

var (
	points   []*point
	field    [][]fieldInfo
	infinite []*point
)

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func manhattan(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func isInfinite(p *point) bool {
	for _, v := range infinite {
		if v == p {
			return true
		}
	}
	return false
}

func Star1() {
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
		for j := 0; j < maxY; j++ {
			field[i][j].minDistance = MIN_DST_UNDEFINED
		}
	}

	// setup starting points
	for _, v := range points {
		field[v.x][v.y].closestCoordinates = []*point{v}
		field[v.x][v.y].minDistance = MIN_DST_STARTING_POINT
	}

	// calculate distances for each cell in field
	for pIdx, p := range points {
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				dst := manhattan(point{x, y}, *p)
				if dst < field[x][y].minDistance {
					field[x][y].minDistance = dst
					field[x][y].closestCoordinates = []*point{points[pIdx]}
				} else if dst == field[x][y].minDistance {
					field[x][y].closestCoordinates = append(field[x][y].closestCoordinates, p)
				}
			}
		}
	}

	// gather points wiht infinite area
	infinite = make([]*point, 0)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			f := field[x][y]
			if len(f.closestCoordinates) > 1 {
				continue
			}
			if x == 0 || y == 0 || x == maxX-1 || y == maxY-1 {
				alreadyAdded := false
				for _, v := range infinite {
					if v == f.closestCoordinates[0] {
						alreadyAdded = true
						break
					}
				}
				if !alreadyAdded {
					infinite = append(infinite, f.closestCoordinates[0])
				}
			}
		}
	}

	// calculate non shared area of each starting point
	area := make(map[*point]int)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			f := field[x][y]
			if len(f.closestCoordinates) > 1 {
				continue
			}
			for _, coords := range f.closestCoordinates {
				area[coords] += 1
			}
		}
	}

	maxArea := 0
	for k, v := range area {
		if !isInfinite(k) && v > maxArea {
			maxArea = v
		}
	}

	fmt.Println("day6 star1:", maxArea)
}

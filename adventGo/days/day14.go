package days

import (
	"log"
	"strconv"
	"strings"
)

type Coord struct {
	x int64
	y int64
}

type Cave struct {
	cave map[Coord]string
	minX int64
	maxX int64
	// away from 0 so maxY is distance from origin down
	minY int64
	maxY int64
}

func findMin[T int64 | int](a, b T) int64 {
	if a < b {
		return int64(a)
	}
	return int64(b)
}

func findMax[T int64 | int](a, b T) int64 {
	if a > b {
		return int64(a)
	}
	return int64(b)
}

func Day14(sb string) {
	tsb := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`
	tsb = tsb
	// p 1
	testcase := tsb
	sandStart := Coord{x: 500, y: 0}
	cave := Cave{minX: sandStart.y, minY: sandStart.y, maxX: sandStart.x, maxY: sandStart.x, cave: make(map[Coord]string, 1)}
	cave.cave[sandStart] = "+"
	log.Printf("====> \n%v %v", testcase, sandStart)

	// draw the cave
	// fill map with blocks and empties can extrapolate drawing later if I need to

	cursor := Coord{x: 0, y: 0}
	for _, row := range strings.Split(testcase, "\n") {
		inputs := strings.Split(row, " -> ")
		for j := 1; j < len(inputs); j++ {
			// have range now
			start := strings.Split(inputs[j-1], ",")
			end := strings.Split(inputs[j], ",")
			x0, _ := strconv.ParseInt(start[0], 10, 0)
			y0, _ := strconv.ParseInt(start[1], 10, 0)
			x1, _ := strconv.ParseInt(end[0], 10, 0)
			y1, _ := strconv.ParseInt(end[1], 10, 0)

			newCoords := make(map[Coord]string, 0)
			var min, max int64
			min, max = 0, 0
			if x0-x1 != 0 {
				min = findMin(x0, x1)
				max = findMax(x0, x1)
				for k := min; min <= max; k++ {
					newCoords[Coord{x: k, y: y0}] = "x"
				}

			} else if y0-y1 != 0 {
				min = findMin(y0, y1)
				max = findMax(y0, y1)
				for k := min; min <= max; k++ {
					newCoords[Coord{x: y0, y: k}] = "x"
				}
			}

			cave.cave[cursor] = "x"
		}
		log.Print(cave.cave)
	}

	// do while loop for tick

	// if one sand grain is in motion, dont start another
}

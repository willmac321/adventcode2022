package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
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

func printCave(cave map[Coord]string) {
	keys := maps.Keys(cave)
	s := ""
	for r := 0; r < 15; r++ {
		s += strconv.Itoa(r)
		if r < 10 {
			s += " "
		}
		for i := 480; i < 515; i++ {
			isMatch := false
			for _, k := range keys {
				if i == int(k.x) && r == int(k.y) {
					isMatch = true
					s += cave[k]
				}
			}
			if !isMatch {
				s += "."
			}
		}
		s += "\n"
	}
	fmt.Println(s)
}

func Day14(sb string) {
	tsb := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`
	tsb = tsb
	// p 1
	testcase := sb
	sandStart := Coord{x: 500, y: 0}
	cave := Cave{minX: sandStart.x, minY: sandStart.y, maxX: sandStart.x, maxY: sandStart.y, cave: map[Coord]string{sandStart: "+"}}
	cave.cave[sandStart] = "+"

	// draw the cave
	// fill map with blocks and empties can extrapolate drawing later if I need to

	for _, row := range strings.Split(testcase, "\n") {
		if row == "" {
			continue
		}
		inputs := strings.Split(row, " -> ")
		for j := 1; j < len(inputs); j++ {
			// have range now
			start := strings.Split(inputs[j-1], ",")
			end := strings.Split(inputs[j], ",")
			x0, _ := strconv.ParseInt(start[0], 10, 0)
			y0, _ := strconv.ParseInt(start[1], 10, 0)
			x1, _ := strconv.ParseInt(end[0], 10, 0)
			y1, _ := strconv.ParseInt(end[1], 10, 0)

			var min, max int64
			if x0-x1 != 0 {
				min = findMin(x0, x1)
				max = findMax(x0, x1)
				// set cave max min
				if min < cave.minX {
					cave.minX = min
				}
				if max > cave.maxX {
					cave.maxX = max
				}
				if y0 < cave.minY {
					cave.minY = y0
				}
				if y0 > cave.maxY {
					cave.maxY = y0
				}
				for k := min; k <= max; k++ {
					cave.cave[Coord{x: k, y: y0}] = "x"
				}
			} else if y0-y1 != 0 {
				min = findMin(y0, y1)
				max = findMax(y0, y1)
				if y1 > 149 || y0 > 149 {
				}
				// set cave max min
				if x0 < cave.minX {
					cave.minX = x0
				}
				if x0 > cave.maxX {
					cave.maxX = x0
				}
				if min < cave.minY {
					cave.minY = min
				}
				if max > cave.maxY {
					cave.maxY = max
				}
				for k := min; k <= max; k++ {
					cave.cave[Coord{x: x0, y: k}] = "x"
				}
			}
		}
	}
	printCave(cave.cave)
	log.Printf("maxY: %v minY: %v", cave.maxY, cave.minY)

	/// part 2 - find floor

	cave.maxY += 2

	for i := cave.minX - 1000; i < cave.maxX+1000; i++ {
		cave.cave[Coord{x: i, y: cave.maxY}] = "x"
	}

	/// -------------------

	// start sand and do while loop for tick
	// if one sand grain is in motion, dont start another
	sandStopCount := 0
	sandFlow := true
	sandGrain := Coord{x: 500, y: 0}
	for sandFlow {
		// if a single grain has stopped
		hasStopped := false
		// fall directly down if possible
		if _, exists := cave.cave[Coord{x: sandGrain.x, y: sandGrain.y + 1}]; !exists {
			sandGrain = Coord{x: sandGrain.x, y: sandGrain.y + 1}
		} else if _, exists := cave.cave[Coord{x: sandGrain.x - 1, y: sandGrain.y + 1}]; !exists {
			sandGrain = Coord{x: sandGrain.x - 1, y: sandGrain.y + 1}
		} else if _, exists := cave.cave[Coord{x: sandGrain.x + 1, y: sandGrain.y + 1}]; !exists {
			sandGrain = Coord{x: sandGrain.x + 1, y: sandGrain.y + 1}
		} else {
			hasStopped = true
			if sandGrain.y == 0 {
				sandFlow = false
			}
		}

		if hasStopped {
			sandStopCount++
			cave.cave[sandGrain] = "o"
			sandGrain = Coord{x: 500, y: 0}
		}

		// stop condition, sand has fallen past cave bottom
		if sandGrain.y > cave.maxY {
			sandFlow = false
		}
	}
	printCave(cave.cave)
	log.Printf("grains: %v", sandStopCount)
}

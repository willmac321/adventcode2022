package day15

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Coord struct {
	x int64
	y int64
}

type Exclude struct {
	x0 int64
	x1 int64
}

func manDist(x0, y0, x1, y1 int64) int64 {
	return int64(math.Abs(float64(x0)-float64(x1))) + int64(math.Abs(float64(y0)-float64(y1)))
}

func printGrid(m map[Coord]bool, row int) {
	keys := maps.Keys(m)
	s := ""
	start, stop := -50, 50
	for r := row - 1; r < row+1; r++ {
		s += strconv.Itoa(r)
		if r < 10 && r >= 0 {
			s += " "
		}
		for i := start; i < stop; i++ {
			isMatch := false
			for _, k := range keys {
				if i == int(k.x) && r == int(k.y) {
					isMatch = true
					s += "#"
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

func Day15P1(sb string) {
	// tsb1 := ` Sensor at x=8, y=7: closest beacon is at x=2, y=10
	// `
	tsb := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

	var cutoff int64
	cutoff = 2000000
	tsb = tsb
	// p 1
	testcase := sb
	re := regexp.MustCompile(`-?[0-9]+`)

	grid := make(map[Coord]bool)

	for _, input := range strings.Split(testcase, "\n") {
		if input == "" {
			continue
		}
		coords := re.FindAllString(input, -1)
		sX, _ := strconv.ParseInt(coords[0], 10, 64)
		sY, _ := strconv.ParseInt(coords[1], 10, 64)
		bX, _ := strconv.ParseInt(coords[2], 10, 64)
		bY, _ := strconv.ParseInt(coords[3], 10, 64)

		manD := manDist(sX, sY, bX, bY)
		// minX := sX - manD
		// maxX := sX + manD
		minY := sY - manD
		maxY := sY + manD
		log.Printf("%v manD: %v", input, manD)
		if cutoff < minY || cutoff > maxY {
			continue
		}
		for row := minY; row <= maxY; row++ {
			if row != cutoff {
				continue
			}
			xW := 2*manD - (2*sY - 2*row)
			if row > sY {
				xW = 2*manD + (2*(sY-manD) - 2*(row-manD))
			}
			for col := sX - xW/2; col <= sX+xW/2; col++ {
				if _, exists := grid[Coord{x: col, y: row}]; !exists {
					if col == sX && row == sY {
						grid[Coord{x: col, y: row}] = false
					} else if col == bX && row == bY {
						grid[Coord{x: col, y: row}] = false
					} else {
						grid[Coord{x: col, y: row}] = true
					}
				}
			}
		}
	}
	printGrid(grid, int(cutoff))
	count := 0
	for _, v := range maps.Values(grid) {
		if v {
			count++
		}
	}
	log.Printf("vals in that row:%v", count)
}

type byLength []Exclude

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return s[i].x0 < s[j].x0
}

// p2
func Day15(sb string) {
	// tsb1 := ` Sensor at x=8, y=7: closest beacon is at x=2, y=10
	// `
	tsb := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

	var cutoffM int64
	cutoffM = 4000000
	tsb = tsb
	// p 1
	testcase := sb
  re := regexp.MustCompile(`-?[0-9]+`)

	for i := 0; i < int(cutoffM)+1; i++ {
		grid := make([]Exclude, 0)
		cutoff := int64(i)

		for _, input := range strings.Split(testcase, "\n") {
			if input == "" {
				continue
			}
			coords := re.FindAllString(input, -1)
			sX, _ := strconv.ParseInt(coords[0], 10, 64)
			sY, _ := strconv.ParseInt(coords[1], 10, 64)
			bX, _ := strconv.ParseInt(coords[2], 10, 64)
			bY, _ := strconv.ParseInt(coords[3], 10, 64)

			manD := manDist(sX, sY, bX, bY)
			row := cutoff
			// find the width
			xW := 2*manD - (2*sY - 2*row)
			if row > sY {
				xW = 2*manD + (2*(sY-manD) - 2*(row-manD))
			}
			if xW <= 0 {
				continue
			}

			var colMin int64
			var colMax int64
			colMin = 0
			colMax = cutoffM

			if d := sX - xW/2; d > colMin {
				colMin = d
			}
			if d := sX + xW/2; d < colMax {
				colMax = d
			}

			if colMax < colMin {
				temp := colMin
				colMin = colMax
				colMax = temp
			}
			grid = append(grid, Exclude{x0: colMin, x1: colMax})
		}

		sort.Sort(byLength(grid))
		var curr Exclude
		curr, grid = grid[0], grid[1:]
		for len(grid) > 0 {
			var g Exclude
			g, grid = grid[0], grid[1:]
			if g.x1 > curr.x1 && g.x0 <= curr.x1+1 {
				curr.x1 = g.x1
			}
			if g.x0 < curr.x0 && g.x1 >= curr.x0-1 {
				curr.x0 = g.x0
			}
		}

		if curr.x0 > 0 || curr.x1 < cutoffM {
			log.Printf("found it: (%v, %v)" , curr.x1 + 1, cutoff)
			log.Printf("math: %v" , (curr.x1 +1) * cutoffM + cutoff)
			break
		}
		if cutoff%100000 == 0 {
			log.Printf("row number: %v", (curr.x1 +1) * cutoffM + cutoff)
		}
	}
}

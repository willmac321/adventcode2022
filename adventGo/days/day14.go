package days

import (
	"log"
)

type Coord struct {
	x int
	y int
	str "sand" | "rock" | "air"
}

type Cave struct {
	cave map[Coord]string
	minX int
	maxX int
	// away from 0 so maxY is distance from origin down
	minY int
	maxY int
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
	
	for _,input:=range testcase {



	}


	// do while loop for tick

	// if one sand grain is in motion, dont start another
}

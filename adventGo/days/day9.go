package days

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func moveHead(dir string, count int) (x1 int, y1 int) {
	x1, y1 = 0, 0
	switch dir {
	case "R":
		x1 += count
	case "L":
		x1 -= count
	case "U":
		y1 += count
	case "D":
		y1 -= count
	}
	return
}

func moveTail(xH int, yH int, xT0 int, yT0 int, dir string) (xT int, yT int) {
	xT, yT = xT0, yT0
	distx := xH - xT0
	disty := yH - yT0

	// is diagonal?
	switch dir {
	case "U", "D":
		if math.Abs(float64(distx)) > 0 && math.Abs(float64(disty)) > 1 {
			if distx < 0 {
				xT--
			} else {
				xT++
			}
		}
	case "R", "L":
		if math.Abs(float64(disty)) > 0 && math.Abs(float64(distx)) > 1 {
			if disty < 0 {
				yT--
			} else {
				yT++
			}
		}

	}

	if math.Abs(float64(distx)) > 1 {
		if distx < 0 {
			xT--
		} else {
			xT++
		}
	}
	if math.Abs(float64(disty)) > 1 {
		if disty < 0 {
			yT--
		} else {
			yT++
		}
	}

	return
}

func moveHeadAndTail(xH0 int, yH0 int, xT0 int, yT0 int, dir string, count int, tailMap0 map[string]bool) (xH int, yH int, xT int, yT int, tailMap map[string]bool) {
	tailMap = tailMap0
	xH, yH, xT, yT = xH0, yH0, xT0, yT0
	i := 0

	for i < count {
		xD, yD := moveHead(dir, 1)
		xH += xD
		yH += yD
		xT, yT = moveTail(xH, yH, xT, yT, dir)
		str := strconv.Itoa(xT) + "," + strconv.Itoa(yT)
		tailMap[str] = true
		i++
	}
	return
}

func Day9(sb string) {
	tsb := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
	tsb2 := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
	tsb = tsb
	tsb2 = tsb2
	input := strings.Split(tsb2, "\n")
	input = input[:len(input)-1]

	xH, yH, xT, yT := 0, 0, 0, 0

	tailMap := make(map[string]bool)
	tailMap["0,0"] = true
	for _, d := range input {
		direction := strings.Split(d, " ")
		count, err := strconv.ParseInt(direction[1], 10, 0)
		if err != nil {
			log.Fatalln(err)
		}
		xH, yH, xT, yT, tailMap = moveHeadAndTail(xH, yH, xT, yT, direction[0], int(count), tailMap)
	}
	output1 := 0
	for range tailMap {
		output1++
	}
	log.Printf("Tail visits %v", output1)
	Day9_p2(sb)
}

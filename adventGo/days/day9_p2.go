package days

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Segment struct {
	x int
	y int
}

func moveHeadAndTails[T []Segment](t2d T, dir string, count int, tailMap0 map[string]bool) (t2 T, tailMap map[string]bool) {
	t2 = t2d
	tailMap = tailMap0
	i := 0
	for i < count {
		// move the head
		xD, yD := moveHead(dir, 1)
		t2[0].x += xD
		t2[0].y += yD

		for j := 1; j < len(t2); j++ {
			distx := t2[j-1].x - t2[j].x
			disty := t2[j-1].y - t2[j].y

			if math.Abs(float64(distx)) > 0 && math.Abs(float64(disty)) > 1 {
				// diagnol max y
				if distx < 0 {
					t2[j].x--
				} else {
					t2[j].x++
				}
				if disty < 0 {
					t2[j].y--
				} else {
					t2[j].y++
				}
			} else if math.Abs(float64(disty)) > 0 && math.Abs(float64(distx)) > 1 {
				if distx < 0 {
					t2[j].x--
				} else {
					t2[j].x++
				}
				if disty < 0 {
					t2[j].y--
				} else {
					t2[j].y++
				}
			} else if math.Abs(float64(distx)) > 1 {
				if distx < 0 {
					t2[j].x--
				} else {
					t2[j].x++
				}
			} else if math.Abs(float64(disty)) > 1 {
				if disty < 0 {
					t2[j].y--
				} else {
					t2[j].y++
				}
			}
			if j == 9 {
				str := strconv.Itoa(t2[j].x) + "," + strconv.Itoa(t2[j].y)
				tailMap[str] = true
			}
		}

		i++
	}

	return
}

func printLine(t2 []Segment, direction []string) {
	out := ""
	y := 21
	for y >= -13 {
		x := -13
		for x < 13 {
			t := "."
			for i := range t2 {
				if y == t2[i].y && x == t2[i].x {
					t = strconv.Itoa(i)
				}
			}
			out += t
			x++
		}
		out += "\n"
		y--
	}
	log.Println(direction)
	fmt.Println(out)
}

func printTail(tail map[string]bool) {
	out := ""
	y := 50
	tailArr := maps.Keys(tail)
	for y >= -51 {
		x := -50
		for x < 51 {
			t := "."
			for _, v := range tailArr {
				x0, _ := strconv.ParseInt(strings.Split(v, ",")[0], 10, 0)
				y0, _ := strconv.ParseInt(strings.Split(v, ",")[1], 10, 0)
				if y == int(y0) && x == int(x0) {
					t = "x"
				}
			}
			out += t
			x++
		}
		out += "\n"
		y--
	}
	fmt.Println(out)
}

func Day9_p2(sb string) {
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
	input := strings.Split(sb, "\n")
	input = input[:len(input)-1]

	tailMap := make(map[string]bool)
	tailMap["0,0"] = true
	start := Segment{x: 0, y: 0}
	// include head
	t2 := make([]Segment, 10)

	for j := range t2 {
		t2[j] = start
	}

	for _, d := range input {
		direction := strings.Split(d, " ")
		count, err := strconv.ParseInt(direction[1], 10, 0)
		if err != nil {
			log.Fatalln(err)
		}
		t2, tailMap = moveHeadAndTails(t2, direction[0], int(count), tailMap)
		// printLine(t2, direction)
	}
	moves:= len(maps.Values(tailMap))
	log.Printf("LongBoy tail move: %v\n", moves)
	printTail(tailMap)
}

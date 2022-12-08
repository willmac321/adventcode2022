package days

import (
	"log"
	"strconv"
	"strings"
)

func isVisLeft(column int, row []int, x int) bool {
	for i := x - 1; i > -1; i-- {
		if row[i] >= column {
			return false
		}
	}
	return true
}

func isVisRight(column int, row []int, x int) bool {
	for i := x + 1; i < len(row); i++ {
		if row[i] >= column {
			return false
		}
	}
	return true
}

func isVisUp(column int, grid [][]int, x int, y int) bool {
	for j := y - 1; j > -1; j-- {
		if grid[j][x] >= column {
			return false
		}
	}
	return true
}

func isVisDown(column int, grid [][]int, x int, y int) bool {
	for j := y + 1; j < len(grid); j++ {
		if grid[j][x] >= column {
			return false
		}
	}
	return true
}

func countVisible(grid [][]bool) (rv int) {
	rv = 0
	for _, row := range grid {
		for _, item := range row {
			if item {
				rv++
			}
		}
	}
	return
}

// p---------------p2

func countVisLeft(column int, row []int, x int) (rv int) {
	rv = 0
	for i := x - 1; i > -1; i-- {
		rv++
		// if cant see past tree
		if row[i] >= column {
			return rv
		}
	}
	return
}

func countVisRight(column int, row []int, x int) (rv int){
	rv = 0
	for i := x + 1; i < len(row); i++ {
		rv++
		if row[i] >= column {
			return rv
		}
	}
	return
}

func countVisUp(column int, grid [][]int, x int, y int) (rv int) {
	rv = 0
	for j := y - 1; j > -1; j-- {
		rv++
		if grid[j][x] >= column {
			return rv
		}
	}
	return
}

func countVisDown(column int, grid [][]int, x int, y int) (rv int){
	rv = 0
	for j := y + 1; j < len(grid); j++ {
		rv++
		if grid[j][x] >= column {
			return rv
		}
	}
	return
}
// ------------------

func Day8(sb string) {
	tsb := `30373
25512
65332
33549
35390`
	tsb = tsb
	input := make([][]int, 0)
	for _, o := range strings.Split(sb, "\n") {
		if o != "" {
			t := strings.Split(o, "")
			row := make([]int, 0)
			for _, r := range t {
				out, err := strconv.ParseInt(r, 10, 0)
				if err != nil {
					log.Panic(err)
				}
				row = append(row, int(out))
			}
			input = append(input, row)
		}
	}
	output1 := make([][]bool, 0)
	output2 := 0
	for y, row := range input {
		rowOut := make([]bool, 0)
		for x, column := range row {
			rowOut = append(rowOut, isVisLeft(column, row, x) || isVisRight(column, row, x) || isVisUp(column, input, x, y) || isVisDown(column, input, x, y))
			if t := countVisLeft(column, row, x) * countVisRight(column, row, x) * countVisUp(column, input, x, y) * countVisDown(column, input, x, y); t > output2 {
				output2 = t
			}
		}
		output1 = append(output1, rowOut)
	}
	log.Printf("%v", countVisible(output1))
	log.Printf("%v", output2)
}

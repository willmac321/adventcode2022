package days

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Day10(sb string) {
	tsb := `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`
	tsb = tsb
	input := strings.Split(sb, "\n")
	input = input[:len(input)-1]
	day10P1(input)
	day10P2(input)
}

func performOp(input []string, xReg0 int, signalStrength0 int, cycles0 int) (xReg int, signalStrength int, cycles int) {
	xReg = xReg0
	cycles = cycles0 + 1
	signalStrength = signalStrength0
	if cycles%20 == 0 && math.Mod(float64(cycles)/20, 2) != 0 {
		signalStrength += (cycles * xReg)
		log.Printf("%v %v", cycles, (cycles * xReg))
	}
	switch input[0] {
	case "noop":
		// no op
	case "addx":
		val, err := strconv.ParseInt(input[1], 10, 0)
		if err != nil {
			log.Fatalln(err)
		}
		// get the first cycle in the 2 cycle addx
		xReg, signalStrength, cycles = performOp([]string{"noop"}, xReg, signalStrength, cycles)
		xReg += int(val)
	}
	return
}

func day10P1(inputs []string) {
	xReg, signalStrength, cycles := 1, 0, 0
	for _, input := range inputs {
		t := strings.Split(input, " ")
		xReg, signalStrength, cycles = performOp(t, xReg, signalStrength, cycles)
	}
	log.Printf("total cycles: %v total signalStrength: %v", cycles, signalStrength)
}

func performOp2(input []string, xReg0 int, rowOut0 string, cycles0 int) (xReg int, strOut string, cycles int) {
	xReg = xReg0
	cycles = cycles0 + 1
	strOut = rowOut0
	pointer := cycles - ((cycles - 1) / 40 * 40) - 1
	if pointer == xReg || pointer+1 == xReg || pointer-1 == xReg {
		strOut += "#"
	} else {
		strOut += "."
	}
	if cycles%40 == 0 {
		strOut += "\n"
	}
	switch input[0] {
	case "noop":
		// no op
	case "addx":
		val, err := strconv.ParseInt(input[1], 10, 0)
		if err != nil {
			log.Fatalln(err)
		}
		// get the first cycle in the 2 cycle addx
		xReg, strOut, cycles = performOp2([]string{"noop"}, xReg, strOut, cycles)
		xReg += int(val)
	}
	return
}

func day10P2(inputs []string) {
	xReg, signalStrength, cycles := 1, "", 0
	for _, input := range inputs {
		t := strings.Split(input, " ")
		xReg, signalStrength, cycles = performOp2(t, xReg, signalStrength, cycles)
	}
	fmt.Printf(signalStrength)
}

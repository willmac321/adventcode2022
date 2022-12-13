package days

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	stack []int64
}

func doOp(item1 int64, item2 int64, operand string) int64 {
	var rv int64
	switch operand {
	case "*":
		rv = item1 * item2
	case "/":
		rv = item1 / item2
	case "+":
		rv = item1 + item2
	case "-":
		rv = item1 - item2
	case "%":
		rv = item1 % item2
	}
	return rv
}

func Day11(sb string) {
	tsb := `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`
	tsb = tsb
	input := strings.Split(sb, "\n\n")
	day11P1(input, false)
	day11P2(input)
}

func day11P1(inputs []string, isPart2 bool) {
	monkeyGang := make(map[string]Monkey)
	rounds := 10000
	round := 0
	monkeyCount := make(map[string]int64)

	var testMod int64
	testMod = 1
	for _, stringinput := range inputs {
		input := strings.Split(stringinput, "\n")
		rColon, _ := regexp.Compile(": ")
		test := rColon.Split(input[3], 2)[1]
		strTestDivisor := strings.Split(test, "by ")[1]
		testDivisor, _ := strconv.ParseInt(strTestDivisor, 10, 64)
		testMod *= testDivisor
	}
	for round < rounds {
		for _, stringinput := range inputs {

			// get monkey

			input := strings.Split(stringinput, "\n")
			r, _ := regexp.Compile("[0-9]+")
			rColon, _ := regexp.Compile(": ")

			monkeyNumber := r.FindString(input[0])
			start := r.FindAllString(input[1], 25)
			worry := make([]int64, len(start))
			// turn to int
			for i, v := range start {
				c, _ := strconv.ParseInt(v, 10, 64)
				worry[i] = c
			}

			operation := rColon.Split(input[2], 2)[1]
			test := rColon.Split(input[3], 2)[1]
			testTrue := rColon.Split(input[4], 2)[1]
			testFalse := rColon.Split(input[5], 2)[1]

			// put monkey items on monkey stack
			if round == 0 {
				if monke, exists := monkeyGang[monkeyNumber]; exists {
					monke.stack = append(worry, monke.stack...)
					monkeyGang[monkeyNumber] = monke
				} else {
					monke.stack = worry
					monkeyGang[monkeyNumber] = monke
				}
			}

			// monkey do
			monkeDo := true
			if len(monkeyGang[monkeyNumber].stack) < 1 {
				monkeDo = false
			}

			for monkeDo {
				// get 0 index el in array
				monke := monkeyGang[monkeyNumber]
				stack := monke.stack
				var item int64
				if len(stack) > 1 {
					item, stack = stack[0], stack[1:]
				} else {
					item, stack = stack[0], []int64{}
				}
				monke.stack = stack
				monkeyGang[monkeyNumber] = monke

				// count monke
				monkeyCount[monkeyNumber]++

				// do operation
				rOld, _ := regexp.Compile("old")
				requal, _ := regexp.Compile("= ")
				operationLocal := rOld.ReplaceAllString(operation, strconv.FormatInt(item, 10))
				operationLocal = requal.Split(operationLocal, 2)[1]
				op := strings.Split(operationLocal, " ")
				item1, _ := strconv.ParseInt(op[0], 10, 64)
				item2, _ := strconv.ParseInt(op[2], 10, 64)
				output := doOp(item1, item2, op[1])

				// do reduce
				if !isPart2 {
					output = output / 3
				} else {
					output = output % testMod
				}

				// do test
				strTestDivisor := strings.Split(test, "by ")[1]
				testDiv, _ := strconv.ParseInt(strTestDivisor, 10, 64)
				testOut := doOp(output, testDiv, "%") == 0

				var rMonke string
				if testOut {
					rMonke = strings.Split(testTrue, "monkey ")[1]
				} else {
					rMonke = strings.Split(testFalse, "monkey ")[1]
				}
				if monke, exists := monkeyGang[rMonke]; exists {
					monke.stack = append(monke.stack, output)
					monkeyGang[rMonke] = monke
				} else {
					monke.stack = []int64{output}
					monkeyGang[rMonke] = monke
				}

				if len(monkeyGang[monkeyNumber].stack) < 1 {
					monkeDo = false
				}
			}
		}
		round++
	}
	max := int64(0)
	max2 := int64(0)
	for i := range monkeyCount {
		if max < monkeyCount[i] {
			max = monkeyCount[i]
		}
	}
	for i := range monkeyCount {
		if max2 < monkeyCount[i] && max != monkeyCount[i] {
			max2 = monkeyCount[i]
		}
	}

	log.Printf("%v\n monkey business:%v", monkeyCount, max*max2)
}

func day11P2(inputs []string) {
	day11P1(inputs, true)
}

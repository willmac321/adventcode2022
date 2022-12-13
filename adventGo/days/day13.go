package days

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type List struct {
	val      int
	children []*List
}

func isNumber(st string) bool {
	numb, _ := regexp.Compile("[0-9]+")
	bracketL, _ := regexp.Compile(`\[`)
	bracketR, _ := regexp.Compile(`\]`)
	return !bracketL.MatchString(st) && !bracketR.MatchString(st) && numb.MatchString(st)
}

func isLeftBrack(st string) bool {
	bracketL, _ := regexp.Compile(`\[`)
	return bracketL.MatchString(st)
}

func isRightBrack(st string) bool {
	bracketR, _ := regexp.Compile(`\]`)
	return bracketR.MatchString(st)
}

func findClose(st string) int {
	// find closing bracket index
	brackC := 1
	closingI := -99
	for j := 0; j < len(st); j++ {
		if string(st[j]) == "[" {
			brackC++
		}
		if string(st[j]) == "]" {
			brackC--
		}
		if brackC == 0 {
			closingI = j
			break
		}
	}
	return closingI
}

func compare(input []string) int {
	left, right := input[0], input[1]
	if len(right) < 1 && len(left) < 1 {
		return 0
	}

	// if empty right abort & fail
	if len(right) < 1 {
		return -1
	}
	if len(left) < 1 {
		return 1
	}

	var leftVal int64
	var rightVal int64
	leftVal, rightVal = -1, -1

	leftStr, rightStr := "", ""
	// shift left and right
	itemL, left, _ := strings.Cut(left, ",")
	itemR, right, _ := strings.Cut(right, ",")

	// if the item is only a number, getnumber
	if isNumber(itemL) {
		leftVal, _ = strconv.ParseInt(itemL, 0, 0)
	}
	if isNumber(itemR) {
		rightVal, _ = strconv.ParseInt(itemR, 0, 0)
	}
	if itemL == "" && isLeftBrack(left) {
		itemL = left[1:]
	}
	if itemR == "" && isLeftBrack(right) {
		itemR = right[1:]
	}

	if itemR != "" && string(itemR[0])==","{
		itemR=itemR[1:]
	}
	if itemL != "" && string(itemL[0])==","{
		itemL=itemL[1:]
	}


	if isLeftBrack(itemL) {
		itemL = itemL[1:] + "," + left
		endB := -99
		// find closing bracket index
		endB = findClose(itemL)
		// grab that subst
		log.Printf("input %v, itemL: %v, leftVal: %v, leftStr: %v, left: %v, endB: %v", input[0], itemL, leftVal, leftStr, left, endB)
		leftStr = itemL[:endB]
		left = itemL[endB+1:]
	}

	if isLeftBrack(itemR) {
		itemR = itemR[1:] + "," + right
		endB := -99
		// find closing bracket index
		endB = findClose(itemR)
		// grab that subst
		rightStr = itemR[:endB]
		right = itemR[endB+1:]
	}

	// if both are intscompare and go next
	if leftVal != -1 && rightVal != -1 {
		if leftVal < rightVal {
			// pass if left is smaller
			return 1
		}
		if leftVal > rightVal {
			// fail if left is greater
			return -1
		}
		if leftVal == rightVal {
			// equal, keep going down the rabbit hole
			return compare([]string{left, right})
		}
	}

	// if both are lists
	if rightVal == -1 && leftVal == -1 {
		rem, sub := compare([]string{left, right}), compare([]string{leftStr, rightStr})
		if sub == -1 || rem == -1 {
			return -1
		} else if sub == 1 || rem == 1 {
			return 1
		} else {
			return 0
		}
	}

	// test to see if one is int and one is list, if so convert to int and go
	if (rightVal == -1 || leftVal == -1) && (rightVal != -1 || leftVal == -1) {
		if rightVal == -1 {
			sub := compare([]string{"[" + strconv.FormatInt(leftVal, 10) + "]", rightStr})
			if rem := compare([]string{left, right}); sub == -1 || rem == -1 {
				return -1
			} else if sub == 1 || rem == 1 {
				return 1
			} else {
				return 0
			}
		}
		if leftVal == -1 {
			sub := compare([]string{leftStr, "[" + strconv.FormatInt(leftVal, 10) + "]"})
			if rem := compare([]string{left, right}); sub == -1 || rem == -1 {
				return -1
			} else if sub == 1 || rem == 1 {
				return 1
			} else {
				return 0
			}
		}
	}
	// if empty left and haven't return yet consider pass
	return 1
}

func Day13(sb string) {
	tsb := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`
	tsb = tsb
	testcase := sb
	inputs := strings.Split(testcase, "\n\n")
	output := 0
	for i, d := range inputs {
		t := strings.Split(d, "\n")
		out := compare(t)
		log.Print(t, "====>", out)
		if out > 0 {
			output += i + 1
		}
	}
	log.Printf("out ====> %v", output)
}

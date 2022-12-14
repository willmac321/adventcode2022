package days

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
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

func isFirstItemLeftBrack(st string) bool {
	bracketL, _ := regexp.Compile(`^\[`)
	return bracketL.MatchString(st)
}

func splitOnBrack(str string) (string, string) {
	curr := ""
	if isFirstItemLeftBrack(str) {
		str = str[1:]
		rightB := findClose(str)
		// split and get our current compare and the remainder that we'll look at later
		curr = str[:rightB]
		str = str[rightB+1:]
		// if left begins with a comma, drop that because we already split
	}
	if str != "" && string(str[0]) == "," {
		str = str[1:]
	}
	if curr != "" && string(curr[0]) == "," {
		curr = curr[1:]
	}
	return curr, str
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

func createStack(input string) (p []interface{}) {
	json.Unmarshal([]byte(input), &p)
	return
}

func compare(left []interface{}, right []interface{}) int {
	for i, itemLeft := range left {
		if len(right) <= i {
			return -1
		}
		itemRight := right[i]
		vL := reflect.ValueOf(itemLeft)
		vR := reflect.ValueOf(itemRight)

		if vL.Kind() == reflect.Float64 && vR.Kind() == reflect.Float64 {
			if vL.Float() < vR.Float() {
				return 1
			}
			if vL.Float() > vR.Float() {
				return -1
			}
		}

		if vL.Kind() == reflect.Float64 {
			tempL := make([]interface{}, 1)
			tempL[0] = vL.Float()
			if tempR, isOk := itemRight.([]interface{}); isOk {
				if temp := compare(tempL, tempR); temp != 0 {
					return temp
				}
			}
		}
		if vR.Kind() == reflect.Float64 {
			tempR := make([]interface{}, 1)
			tempR[0] = vR.Float()
			if tempL, isOk := itemLeft.([]interface{}); isOk {
				if temp := compare(tempL, tempR); temp != 0 {
					return temp
				}
			}
		}
		tempL, isOkL := itemLeft.([]interface{})
		tempR, isOkR := itemRight.([]interface{})
		if temp := compare(tempL, tempR); isOkL && isOkR && temp != 0 {
			return temp
		}

	}

	if len(left) < len(right) {
		return 1
	}
	return 0
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return compare(createStack(s[i]), createStack(s[j])) == 1
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
	newVal := `[[2]]
[[6]]
`
	tsb = tsb
	// p 1
	testcase := sb
	inputs := strings.Split(testcase, "\n\n")
	output := 0
	for i, d := range inputs {
		t := strings.Split(d, "\n")
		left, right := createStack(t[0]), createStack(t[1])
		out := compare(left, right)
		log.Print(t, "====>", out)
		if out > 0 {
			output += i + 1
		}
	}
	log.Printf("out ====> %v", output)
	// p2
	t := testcase + newVal
	t = strings.ReplaceAll(t, "\n\n", "\n")
	inputs = strings.Split(t, "\n")
	sort.Sort(byLength(inputs))
	output = 1
	for _, d := range inputs {
		fmt.Println(d)
	}
	for i, d := range inputs {
		if d == "[[2]]" || d == "[[6]]" {
			output *= i
		}
	}

	log.Printf("out P2 ====> %v", output)
}

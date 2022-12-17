package day16

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Node struct {
	flow     int
	name     string
	isOpen   bool
	children []string
}

// for the starting node, add 1 minute traversal and 1 minute if opening valve
// if flow is 0 dont spend the minute opening
// mark the node as open
// travel to it's children and follow same structure
// exit when all children are with flow are open and when time is 30
func traverseMinMax(graph map[string]Node, start string, oldpath []string, oldtime int, oldflow int) (time, flow int, path []string) {
	// if out of time bal
	if oldtime >= 3 || len(oldpath) < 1 {
		return oldtime, oldflow, oldpath
	}

	curr := graph[start]
	time = oldtime
	flow = oldflow
	path = make([]string, 0)

	for k := range oldpath {
		if oldpath[k] != start {
			path = append(path, oldpath[k])
		}
	}

	// incrememnt travel time
	time++

	if time >= 30 {
		return oldtime, flow, path
	}

	if !curr.isOpen && curr.flow > 0 {
		time++
		flow += (30 - time) * curr.flow
		if time >= 30 {
			return oldtime, flow, path
		}
	}

	// always set to open so we dont open again
	curr.isOpen = true
	graph[start] = curr

	var maxPath []string
	maxTime := time
	maxFlow := flow
	maxChild := ""

	for _, child := range curr.children {
		t, f, p := traverseMinMax(graph, child, path, time, flow)
		if f > maxFlow && t <= 30 {
			maxChild = child
			maxPath = p
			maxFlow = f
			maxTime = t
		}
	}

	log.Printf("%v %v %v %v", maxChild, maxFlow, maxTime, maxPath)

	if maxFlow > flow {
		flow = maxFlow
		path = maxPath
		time = maxTime
	}

	for _, child := range curr.children {
		t, f, p := traverseMinMax(graph, child, path, time, flow)
		if f > flow && t <= 30 {
			path = p
			flow = f
			time = t
		}
	}

	return
}

func Day16(sb string) {
	tsb := `Valve AA has flow rate=0; tunnels lead to valves DD
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`
	tsb = tsb
	// p 1
	testcase := tsb
	p1 := make(map[string]Node, 0)
	rName := regexp.MustCompile(`[A-Z]{2}`)
	rFlow := regexp.MustCompile(`[0-9]+`)
	for _, input := range strings.Split(testcase, "\n") {
		if input == "" {
			continue
		}
		names := rName.FindAllString(input, -1)
		flow, _ := strconv.ParseInt(rFlow.FindAllString(input, 1)[0], 10, 0)
		p1[names[0]] = Node{name: names[0], flow: int(flow), children: names[1:], isOpen: false}
	}
	log.Printf("%v", p1)

	notVisited := maps.Keys(p1)
	log.Println(traverseMinMax(p1, "AA", notVisited, 0, 0))
}

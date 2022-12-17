package day16

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	flow     int
	name     string
	isOpen   bool
	children []string
}
type byFlow []Node

func (s byFlow) Len() int {
	return len(s)
}

func (s byFlow) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byFlow) Less(i, j int) bool {
	return s[j].isOpen && s[i].flow > s[j].flow
}

// for the starting node, add 1 minute traversal and 1 minute if opening valve
// if flow is 0 dont spend the minute opening
// mark the node as open
// travel to it's children and follow same structure
// exit when all children are with flow are open and when time is 30
func traverseMinMax(graph map[string]Node, start string, oldvisited []string, oldtime int, oldflow int) (time, flow int, visited []string) {
	cutoffTime := 30

	time = oldtime
	flow = oldflow
	visited = []string{}

	path := []string{start}

	for len(path) > 0 && time <= cutoffTime {
		log.Print(time)
		v := path[len(path)-1]
		curr := graph[v]
		visited = append(visited, v)
		log.Print(path)

		if len(path) > 1 {
			path = path[:len(path)-2]
		} else {
			path = make([]string, 0)
		}
		time++
		if !curr.isOpen && curr.flow > 0 {
			time++
			flow += (cutoffTime - time) * curr.flow
		}

		childs := []Node{}
		if !curr.isOpen {
			curr.isOpen = true
			for _, child := range curr.children {
				if len(curr.children) > 1 && visited[len(visited)-1] == child {
					continue
				}
				childs = append(childs, graph[child])
			}
		}
		if curr.isOpen && len(childs) < 1 {
			for _, child := range curr.children {
				if !graph[child].isOpen {
					childs = append(childs, graph[child])
				}
			}
		}
		sort.Sort(byFlow(childs))

		log.Print(childs)
		for _, child := range childs {
			path = append(path, child.name)
		}


		graph[v] = curr
	}
	return
}

func Day16(sb string) {
	tsb := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
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

	log.Println(traverseMinMax(p1, "AA", make([]string, 0), 0, 0))
}

package days

import (
	"log"
	"math"
	"strconv"
)

func getIndex(width int, x int, y int) int {
	return width*y + x
}

func getXY(width int, i int) (x int, y int) {
	x = i % width
	y = i / width
	return
}

type Path struct {
	vertex int
	weight int
}

type Node struct {
	vertex   int
	value    rune
	children []Path
}

func addNode(input []rune, testI int, node Node) Node {
	inp := input[testI]
	if inp == 'S' {
		inp = 'a'
	}
	if inp == 'E' {
		inp = 'z'
	}
	testInp := node.value
	if testInp == 'S' {
		testInp = 'a'
	}
	if testInp == 'E' {
		testInp = 'z'
	}
	if input[testI] != 'S' && inp-testInp <= 1 {
		path := Path{vertex: testI, weight: 1}
		node.children = append(node.children, path)
	}
	return node
}

func djikstraDist(graph []Node, start int, path []int) ([]int, []int) {
	dist := make([]int, len(graph))
	queue := make([]Node, 0)
	for _, g := range graph {
		dist[g.vertex] = math.MaxInt
		path[g.vertex] = -1
		queue = append(queue, g)
	}

	dist[start] = 0
	for len(queue) > 0 {
		// check dist
		d := math.MaxInt
		i := -1
		var u Node
		// find vertex with min distance
		for j, q := range queue {
			if dist[q.vertex] <= d {
				i = j
				d = dist[q.vertex]
				u = q
			}
		}
		// if no more min, panic
		if i < 0 {
			log.Panic("no min is found, uhoh")
			break
		}
		// remove u from q
		queue = append(queue[:i], queue[i+1:]...)

		// for each neighbor of node u still in queue, look at its distance from root node through node q
		// and if its less than what we have currently, replace the recorded distance for the neighbor node
		for _, neighbor := range u.children {
			alt := dist[u.vertex] + neighbor.weight
			if alt < dist[neighbor.vertex] {
				dist[neighbor.vertex] = alt
				path[neighbor.vertex] = u.vertex
			}
		}

	}
	return dist, path
}

func Day12(sb string) {
	tsb := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`
	tsb = tsb
	testcase := sb
	width := 0
	input := make([]rune, 0)
	for _, s := range testcase {
		if string(s) == "\n" {
			break
		}
		width++
	}
	for _, s := range testcase {
		if string(s) != "\n" {
			input = append(input, s)
		}
	}
	day12P1(input, width)
}

func printToWidth(input []int, width int) {
	s := "\n"
	row := 0
	for i := range input {
		if row != i/width {
			s += "\n"
		}
		row = i / width
		s += " "
		if input[i] < 10 && input[i] > -1 {
			s += "  "
		}
		if input[i] < 100 && input[i] > 9 {
			s += " "
		}
		if input[i] == math.MaxInt64 || input[i] == math.MinInt64 || input[i] == -1 {
			s += "  X"
		} else {
			s += strconv.FormatInt(int64(input[i]), 10)
		}
	}
	log.Print(s)
}

func printDirections(input []int, end int) {
	// out := make([]int, len(input))
}

func day12P1(input []rune, width int) {
	// make graph
	graph := make([]Node, len(input))
	start := -1
	end := -1

	for i, s := range input {
		if s == 'S' {
			start = i
		}
		if s == 'E' {
			end = i
		}
		graph[i] = Node{vertex: i, children: nil, value: s}
		x, y := getXY(width, i)
		// get children
		if up := getIndex(width, x, y-1); y > 0 {
			graph[i] = addNode(input, up, graph[i])
		}
		if _, dY := getXY(width, len(input)-1); y < dY {
			down := getIndex(width, x, y+1)
			graph[i] = addNode(input, down, graph[i])
		}
		if x < width-1 {
			right := getIndex(width, x+1, y)
			graph[i] = addNode(input, right, graph[i])
		}
		if x > 0 {
			left := getIndex(width, x-1, y)
			graph[i] = addNode(input, left, graph[i])
		}
	}
	// log.Print(graph, start, end)
	path := make([]int, len(graph))
	var dist []int

	dist, path = (djikstraDist(graph, start, path))
	// get dist to E

	printToWidth(dist, width)
	printToWidth(path, width)
	printDirections(path, end)
	log.Printf("dist from S to end: %v", dist[end])
	// find min dist from a to end
	min := math.MaxInt64
	for i, d := range input {
		path2 := make([]int, len(graph))
		var dist2 []int
		if d == 'a' {
			dist2, path2 = (djikstraDist(graph, i, path2))
			if dist2[end] < min && dist2[end] > 0 {
				min = dist2[end]
			}
		}
	}
	log.Printf("dist from closest a to end: %v", min)
}

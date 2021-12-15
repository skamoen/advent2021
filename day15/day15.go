package day15

import (
	"bufio"
	"fmt"
	"github.com/skamoen/advent2021/util"
	"log"
	"math"
	"os"
	"strconv"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

var cave [][]*node

func (*d) Run() (int, int) {
	//file, err := os.Open("./day15/example.txt")
	file, err := os.Open("./day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var partOneNodes []*node
	var partTwoNodes []*node

	column := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		currentLine := make([]*node, len(line))
		for i := range line {
			c, _ := strconv.Atoi(string(line[i]))
			newNode := &node{
				cost:     c,
				visited:  false,
				distance: math.MaxInt,
				column:   column,
				row:      i,
			}
			currentLine[i] = newNode
			partOneNodes = append(partOneNodes, newNode)
		}
		cave = append(cave, currentLine)
		column++
	}

	start := cave[0][0]
	start.distance = 0

	end := partOneNodes[len(partOneNodes)-1]

	start.visitNeighbors(cave)

	for i := 0; !end.visited; i++ {
		nextNode := nextNode(partOneNodes)
		nextNode.visitNeighbors(cave)
	}

	caveTwo := make([][]*node, len(cave)*5)
	for i := range caveTwo {
		caveTwo[i] = make([]*node, len(cave)*5)
	}
	mX := len(cave)
	mY := len(cave[0])

	for i := 0; i < len(caveTwo); i++ {
		for j := 0; j < len(caveTwo); j++ {
			newCost := (cave[i%mY][j%mX].cost+i/mY+j/mX-1)%9 + 1
			newNode := &node{
				cost:     newCost,
				distance: math.MaxInt,
				column:   i,
				row:      j,
			}
			caveTwo[i][j] = newNode
			partTwoNodes = append(partTwoNodes, newNode)
		}
	}

	startTwo := caveTwo[0][0]
	startTwo.distance = 0

	endTwo := partTwoNodes[len(partTwoNodes)-1]

	startTwo.visitNeighbors(caveTwo)

	for i := 0; !endTwo.visited; i++ {
		nextNode := nextNode(partTwoNodes)
		nextNode.visitNeighbors(caveTwo)
	}

	return end.distance, endTwo.distance
}

func nextNode(n []*node) *node {
	next := &node{distance: math.MaxInt}
	for i := range n {
		if n[i].distance < next.distance && !n[i].visited {
			next = n[i]
		}
	}
	return next
}

func (n *node) visitNeighbors(grid [][]*node) {
	c := n.column
	r := n.row
	// up
	next := getNode(c, r-1, grid)
	if next != nil && !next.visited {
		d := n.distance + next.cost
		if d < next.distance {
			next.distance = d
		}
	}

	// right
	next = getNode(c+1, r, grid)
	if next != nil && !next.visited {

		d := n.distance + next.cost
		if d < next.distance {
			next.distance = d
		}
	}

	//down
	next = getNode(c, r+1, grid)
	if next != nil && !next.visited {
		d := n.distance + next.cost
		if d < next.distance {
			next.distance = d
		}
	}

	//left
	next = getNode(c-1, r, grid)
	if next != nil && !next.visited {
		d := n.distance + next.cost
		if d < next.distance {
			next.distance = d
		}
	}

	n.visited = true
}

func getNode(c, r int, grid [][]*node) (n *node) {
	defer func() {
		if r := recover(); r != nil {
			n = nil
		}
	}()
	n = grid[c][r]
	return
}

type node struct {
	cost     int
	visited  bool
	distance int
	column   int
	row      int
}

func print(n [][]*node) {
	for j := 0; j < len(n[0]); j++ {
		for i := 0; i < len(n); i++ {
			if n[i][j].cost < 99 {
				fmt.Print(n[i][j].cost, "")
			} else {
				fmt.Print("X ")
			}
			if n[i][j].cost < 10 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

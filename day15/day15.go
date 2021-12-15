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

var nodes map[int]*node

var origGridSize int
var gridSize int

var nodesArray []*node

func (*d) Run() (int, int) {
	//file, err := os.Open("./day15/example.txt")
	file, err := os.Open("./day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	column := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		currentLine := make([]*node, len(line))
		for row := range line {
			if nodes == nil {
				nodes = make(map[int]*node, len(line))
				origGridSize = len(line)
				gridSize = len(line)
			}
			c, _ := strconv.Atoi(string(line[row]))
			newNode := &node{
				cost:    c,
				visited: false,
				risk:    math.MaxInt,
				column:  column,
				row:     row,
			}
			currentLine[row] = newNode
			nodesArray = append(nodesArray, newNode)
			nodes[(column+1)*gridSize+row] = newNode
		}
		column++
	}

	start := getNode(0, 0)
	start.risk = 0

	end := nodesArray[len(nodesArray)-1]

	start.visitNeighbors()

	for i := 0; !end.visited; i++ {
		nextNode, s := nextNode(nodesArray)
		nextNode.visitNeighbors()
		nodesArray = append(nodesArray[:s], nodesArray[s+1:]...)
	}

	partOne := end.risk

	gridSize = gridSize * 5
	nodes2 := make(map[int]*node, gridSize)
	nodesArray = make([]*node, 0)
	for _, v := range nodes {
		nodes2[(v.column+1)*gridSize+v.row] = v
		nodesArray = append(nodesArray, v)
	}
	nodes = nodes2

	for i := range nodesArray {
		nodesArray[i].visited = false
		nodesArray[i].risk = math.MaxInt
	}

	startTwo := getNode(0, 0)
	startTwo.risk = 0

	endTwo := getNode(gridSize-1, gridSize-1)

	startTwo.visitNeighbors()

	for i := 0; !endTwo.visited; i++ {
		nextNode, s := nextNode(nodesArray)
		nextNode.visitNeighbors()
		nodesArray = append(nodesArray[:s], nodesArray[s+1:]...)
	}

	return partOne, endTwo.risk
}

func nextNode(n []*node) (*node, int) {
	next := &node{risk: math.MaxInt, cost: -1}
	selected := -1
	for i := range n {
		if !n[i].visited && n[i].risk < next.risk {
			next = n[i]
			selected = i
		}
	}
	return next, selected
}

func (n *node) visitNeighbors() {
	var neighbours = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for _, direction := range neighbours {
		next := getNode(n.column+direction[0], n.row+direction[1])
		if next != nil && !next.visited {
			d := n.risk + next.cost
			if d < next.risk {
				next.risk = d
			}
		}
	}
	n.visited = true
}

func getNode(column, row int) (n *node) {
	if column < 0 || row < 0 || column > gridSize-1 || row > gridSize-1 {
		return nil
	}

	n, ok := nodes[(column+1)*gridSize+row]
	if !ok {
		newCost := (getNode(column%origGridSize, row%origGridSize).cost+column/origGridSize+row/origGridSize-1)%9 + 1
		n = &node{
			cost:   newCost,
			risk:   math.MaxInt,
			column: column,
			row:    row,
		}
		nodes[(column+1)*gridSize+row] = n
		nodesArray = append(nodesArray, n)
		return
	}
	return
}

type node struct {
	cost    int
	visited bool
	risk    int
	column  int
	row     int
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

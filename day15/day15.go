package day15

import (
	"bufio"
	"container/heap"
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

var nodes map[int]*Item

var origGridSize int
var gridSize int

var inputFile = "./day15/input.txt"

//var inputFile = "./day15/example.txt"
var pq PriorityQueue

func (*d) Run() (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	column := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for row := range line {
			if nodes == nil {
				nodes = make(map[int]*Item, len(line))
				origGridSize = len(line)
				gridSize = len(line)
				pq = make(PriorityQueue, 0, len(line)*len(line))
				heap.Init(&pq)
			}
			c, _ := strconv.Atoi(string(line[row]))
			newNode := &Item{
				value:    c,
				priority: math.MaxInt,
				index:    0,
				visited:  false,
				column:   column,
				row:      row,
			}
			nodes[(column+1)*gridSize+row] = newNode
			heap.Push(&pq, newNode)
		}
		column++
	}

	start := getNode(0, 0)
	pq.update(start, start.value, 0)
	// Take start off the PQ
	heap.Pop(&pq)

	end := getNode(gridSize-1, gridSize-1)

	start.visitNeighbors()

	for i := 0; !end.visited; i++ {
		nextNode := heap.Pop(&pq).(*Item)
		nextNode.visitNeighbors()
	}

	partOne := end.priority

	gridSize = gridSize * 5
	pq = make(PriorityQueue, 0, gridSize*gridSize)
	heap.Init(&pq)

	nodes2 := make(map[int]*Item, gridSize*gridSize)
	for _, v := range nodes {
		nodes2[(v.column+1)*gridSize+v.row] = v
		heap.Push(&pq, v)

	}
	nodes = nodes2

	startTwo := getNode(0, 0)
	startTwo.priority = 0

	endTwo := getNode(gridSize-1, gridSize-1)

	startTwo.visitNeighbors()

	for i := 0; !endTwo.visited; i++ {
		nextNode := heap.Pop(&pq).(*Item)
		nextNode.visitNeighbors()
	}

	for _, v := range nodes {
		if v.priority == math.MaxInt {
			fmt.Print("max")
		}
	}
	return partOne, endTwo.priority
}

func (n *Item) visitNeighbors() {
	var neighbours = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for _, direction := range neighbours {
		next := getNode(n.column+direction[0], n.row+direction[1])
		if next != nil && !next.visited {
			d := n.priority + next.value
			if d < next.priority {
				//next.risk = d
				pq.update(next, next.value, d)
			}
		}
	}
	n.visited = true
}

func getNode(column, row int) (n *Item) {
	if column < 0 || row < 0 || column > gridSize-1 || row > gridSize-1 {
		return nil
	}

	n, ok := nodes[(column+1)*gridSize+row]
	if !ok {
		newCost := (getNode(column%origGridSize, row%origGridSize).value+column/origGridSize+row/origGridSize-1)%9 + 1
		n = &Item{
			value:    newCost,
			priority: math.MaxInt,
			visited:  false,
			column:   column,
			row:      row,
		}
		nodes[(column+1)*gridSize+row] = n
		heap.Push(&pq, n)
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
	index   int
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

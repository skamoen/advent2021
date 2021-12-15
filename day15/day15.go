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

	var unvisited []*node

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
			unvisited = append(unvisited, newNode)
		}
		cave = append(cave, currentLine)
		column++
	}

	start := cave[0][0]
	start.distance = 0

	end := unvisited[len(unvisited)-1]

	start.visitNeighbors()

	for i := 0; !end.visited; i++ {
		nextNode := nextNode(unvisited)
		if nextNode.column == len(unvisited)-1 && nextNode.row == len(unvisited)-1 {
			fmt.Println(i)
		}
		nextNode.visitNeighbors()
	}
	//print(cave)
	return end.distance, 0
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

func (n *node) visitNeighbors() {
	c := n.column
	r := n.row
	// up
	next := getNode(c, r-1)
	if next != nil && !next.visited {
		d := n.distance + next.cost
		if d < next.distance {
			next.distance = d
		}
	}

	// right
	next = getNode(c+1, r)
	if next != nil && !next.visited {

		d := n.distance + next.cost
		if d < next.distance {
			next.distance = d
		}
	}

	//down
	next = getNode(c, r+1)
	if next != nil && !next.visited {
		d := n.distance + next.cost
		if d < next.distance {
			next.distance = d
		}
	}

	n.visited = true
}

func getNode(c, r int) (n *node) {
	defer func() {
		if r := recover(); r != nil {
			n = nil
		}
	}()
	n = cave[c][r]
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
			if n[i][j].distance < 99 {
				fmt.Print(n[i][j].distance, "\t")
			} else {
				fmt.Print("X\t")
			}
		}
		fmt.Println()
	}
}

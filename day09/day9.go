package day09

import (
	"bufio"
	"fmt"
	"github.com/skamoen/advent2021/util"
	"log"
	"os"
	"sort"
	"strconv"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() {
	file, err := os.Open("./day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var heights [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]int, len(line))
		for i := range line {
			c, _ := strconv.Atoi(string(line[i]))
			row[i] = c
		}
		heights = append(heights, row)
	}

	riskSum := 0
	var sizes []int
	for i := range heights {
		for j := range heights[i] {
			currentHeight := heights[i][j]

			// Check up
			if j > 0 {
				if heights[i][j-1] == currentHeight || util.MinInt(currentHeight, heights[i][j-1]) != currentHeight {
					continue
				}
			}
			// Check right
			if i < len(heights)-1 {
				if heights[i+1][j] == currentHeight || util.MinInt(currentHeight, heights[i+1][j]) != currentHeight {
					continue
				}
			}
			// Check bottom
			if j < len(heights[i])-1 {
				if heights[i][j+1] == currentHeight || util.MinInt(currentHeight, heights[i][j+1]) != currentHeight {
					continue
				}
			}
			// Check left
			if i > 0 {
				if heights[i-1][j] == currentHeight || util.MinInt(currentHeight, heights[i-1][j]) != currentHeight {
					continue
				}
			}
			riskSum += currentHeight + 1
			b := &basin{start: []int{i, j}}
			size := b.findSize(heights)
			sizes = append(sizes, size)
		}
	}

	sort.Ints(sizes)
	ints := sizes[len(sizes)-3:]
	fmt.Println("Riskfactor", riskSum, "Basin size", ints[0]*ints[1]*ints[2])
}

func (b *basin) findSize(heights [][]int) int {
	b.marker = make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		b.marker[i] = make([]bool, len(heights))
	}

	// Check
	b.checkPoint(b.start[0], b.start[1], heights)

	return b.size
}

func (b *basin) checkPoint(i, j int, heights [][]int) {

	if b.marker[i][j] == false {
		b.size++
		b.marker[i][j] = true
	} else {
		return
	}

	// Check left
	if i > 0 {
		if heights[i-1][j] != 9 {
			b.checkPoint(i-1, j, heights)
		}
	}
	// Check up
	if j > 0 {
		if heights[i][j-1] != 9 {
			b.checkPoint(i, j-1, heights)
		}
	}
	// Check right
	if i < len(heights)-1 {
		if heights[i+1][j] != 9 {
			b.checkPoint(i+1, j, heights)
		}
	}
	// Check bottom
	if j < len(heights[i])-1 {
		if heights[i][j+1] != 9 {
			b.checkPoint(i, j+1, heights)
		}
	}
}

type basin struct {
	start  []int
	marker [][]bool // Mark X, Y as included
	size   int
}

package day11

import (
	"bufio"
	"github.com/skamoen/advent2021/util"
	"log"
	"os"
	"strconv"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() (int, int) {
	file, err := os.Open("./day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	og := octogrid{g: make([][]int, 10)}

	scanner := bufio.NewScanner(file)
	li := 0
	for scanner.Scan() {
		line := scanner.Text()
		var currentLine []int
		for i := range line {
			o, _ := strconv.Atoi(string(line[i]))
			currentLine = append(currentLine, o)
		}
		og.g[li] = currentLine
		li++
	}

	og.doSteps(100)
	partOne := og.flashes

	for i := 0; og.fullFlash[len(og.fullFlash)-1] < 100; i++ {
		og.doSteps(1)
	}

	return partOne, len(og.fullFlash)

}

func (o *octogrid) doSteps(n int) {
	for s := 0; s < n; s++ {
		o.stepFlashes = 0
		o.marker = make([][]bool, len(o.g))
		for i := 0; i < len(o.g); i++ {
			o.marker[i] = make([]bool, len(o.g))
		}

		for i := range o.g {
			for j := range o.g[i] {
				o.increaseAndCheckFlash(i, j)
			}
		}
		o.fullFlash = append(o.fullFlash, o.stepFlashes)
	}
}

type octogrid struct {
	g           [][]int
	flashes     int
	marker      [][]bool
	stepFlashes int
	fullFlash   []int
}

func (o *octogrid) increaseAndCheckFlash(i, j int) {
	if o.marker[i][j] {
		return
	}
	o.g[i][j]++
	if o.g[i][j] > 9 {
		o.flashes++
		o.stepFlashes++
		o.marker[i][j] = true
		o.g[i][j] = 0

		// Check up
		if j > 0 {
			//Check up-left
			if i > 0 {
				o.increaseAndCheckFlash(i-1, j-1)
			}
			o.increaseAndCheckFlash(i, j-1)
			//Check up-right
			if i < len(o.g)-1 {
				o.increaseAndCheckFlash(i+1, j-1)
			}
		}

		// Check left
		if i > 0 {
			o.increaseAndCheckFlash(i-1, j)
		}
		// Check right
		if i < len(o.g)-1 {
			o.increaseAndCheckFlash(i+1, j)
		}
		// Check bottom
		if j < len(o.g[i])-1 {
			// Check down-left
			if i > 0 {
				o.increaseAndCheckFlash(i-1, j+1)
			}

			o.increaseAndCheckFlash(i, j+1)

			// Check down-right
			if i < len(o.g)-1 {
				o.increaseAndCheckFlash(i+1, j+1)
			}
		}

	}
}

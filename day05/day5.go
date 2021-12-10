package day05

import (
	"bufio"
	"github.com/skamoen/advent2021/util"
	"log"
	"os"
	"strconv"
	"strings"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() (int, int) {
	file, err := os.Open("./day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []vector
	var diag []vector

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")
		from, to := split[0], split[1]
		v := makeVectorFromString(from, to)
		if string(v.direction[0]) != "d" {
			lines = append(lines, v)
		} else {
			diag = append(diag, v)
		}
	}

	var m [1000][1000]int
	for _, line := range lines {

		x, y := line.start[0], line.start[1]
		for j := 0; j <= line.length; j++ {
			if line.direction == "h" {
				m[x+j][y]++
			} else if line.direction == "v" {
				m[x][y+j]++
			}
		}
	}

	overlap := 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] >= 2 {
				overlap++
			}
		}
	}

	partOne := overlap

	for _, line := range diag {
		x, y := line.start[0], line.start[1]
		for j := 0; j <= line.length; j++ {
			if line.direction == "du" {
				m[x+j][y+j]++
			} else {
				m[x+j][y-j]++
			}
		}
	}

	overlap = 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] >= 2 {
				overlap++
			}
		}
	}

	return partOne, overlap
}

func makeVectorFromString(from, to string) vector {
	fromSplit := strings.Split(from, ",")
	toSplit := strings.Split(to, ",")

	x1, _ := strconv.Atoi(fromSplit[0])
	y1, _ := strconv.Atoi(fromSplit[1])
	x2, _ := strconv.Atoi(toSplit[0])
	y2, _ := strconv.Atoi(toSplit[1])

	return makeVector(x1, y1, x2, y2)
}

func makeVector(x1, y1, x2, y2 int) vector {
	v := vector{}

	if x1 == x2 {
		v.direction = "v"
		if y1 < y2 {
			v.start = [2]int{x1, y1}
			v.length = y2 - y1
		} else {
			v.start = [2]int{x2, y2}
			v.length = y1 - y2
		}
	} else if y1 == y2 {
		v.direction = "h"
		if x1 < x2 {
			v.start = [2]int{x1, y1}
			v.length = x2 - x1
		} else {
			v.start = [2]int{x2, y2}
			v.length = x1 - x2
		}
	} else {
		if x1 < x2 {
			v.start = [2]int{x1, y1}
			v.length = x2 - x1
			if y1 < y2 {
				v.direction = "du"
			} else {
				v.direction = "dd"
			}
		} else {
			v.start = [2]int{x2, y2}
			v.length = x1 - x2
			if y1 < y2 {
				v.direction = "dd"
			} else {
				v.direction = "du"
			}
		}
	}
	return v
}

type vector struct {
	direction string
	start     [2]int
	length    int
}

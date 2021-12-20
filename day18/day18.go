package day18

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

var inputFile = "./day18/input.txt"

//var inputFile = "./day18/example.txt"

func (*d) Run() (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines = make([][]valueNode, 0, 100)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := parseLine(line)
		lines = append(lines, values)
	}

	var addValues = lines[0]
	maxMag := 0
	for i := 1; i <= len(lines); i++ {
		if i < len(lines) {
			addValues = add(addValues, lines[i])
			addValues = reduce(addValues)
		}

		for j := 1; j <= len(lines); j++ {
			if i != j {
				addA := add(lines[i-1], lines[j-1])
				reduceA := reduce(addA)
				magA := magnitude(reduceA)

				addB := add(lines[j-1], lines[i-1])
				reduceB := reduce(addB)
				magB := magnitude(reduceB)

				if magA > maxMag {
					maxMag = magA
				}
				if magB > maxMag {
					maxMag = magB
				}
			}
		}
	}

	return magnitude(addValues), maxMag
}

func magnitude(values []valueNode) (n int) {
	for len(values) > 1 {
		for i := range values {
			if values[i].depth == values[i+1].depth {
				// merge
				newValue := 3*values[i].value + 2*values[i+1].value
				values[i+1].value = newValue
				values[i+1].depth--
				values = append(values[:i], values[i+1:]...)
				break
			}
		}
	}
	return values[0].value
}

func add(a, b []valueNode) []valueNode {
	var v = make([]valueNode, len(a)+len(b))
	copy(v[:len(a)], a)
	copy(v[len(a):], b)
	for i := range v {
		v[i].depth++
	}
	return v
}

func parseLine(line string) (values []valueNode) {
	depth := 0
	for i := 0; i < len(line); i++ {
		switch string(line[i]) {
		case "[":
			depth++
		case ",":
			continue
		case "]":
			depth--
		default:
			nextComma := strings.Index(line[i:], ",")
			nextClose := strings.Index(line[i:], "]")

			if nextClose == -1 && nextComma == -1 {
				continue
			}

			var valueUntil int
			if nextClose < nextComma || nextComma == -1 {
				valueUntil = nextClose
			} else {
				valueUntil = nextComma
			}

			v, _ := strconv.Atoi(line[i : i+valueUntil])
			values = append(values, valueNode{depth, v})
			i += valueUntil - 1

		}
	}
	return
}

func reduce(values []valueNode) []valueNode {
	for {
		if len(values) > 2 {
			var explode bool
			explode, values = tryExplode(values)
			if explode {
				continue
			}
			var split bool
			split, values = trySplit(values)
			if split {
				continue
			}
			return values
		} else {
			return values
		}
	}
}

func trySplit(values []valueNode) (bool, []valueNode) {
	for i := range values {
		if values[i].value > 9 {
			left := values[i].value / 2
			right := values[i].value - left

			values[i].value = left
			values[i].depth++
			newNode := valueNode{values[i].depth, right}
			arrayEnd := append(
				[]valueNode{newNode},
				values[i+1:]...,
			)
			values = append(
				values[:i+1],
				arrayEnd...,
			)
			return true, values
		}
	}
	return false, values

}

func tryExplode(values []valueNode) (bool, []valueNode) {
	for i := range values {
		if values[i].depth > 4 && values[i+1].depth == values[i].depth {
			goLeft := values[i].value
			if i > 0 {
				values[i-1].value += goLeft
			}

			goRight := values[i+1].value
			if i < len(values)-2 {
				values[i+2].value += goRight
			}

			values[i+1].depth--
			values[i+1].value = 0

			values = append(values[:i], values[i+1:]...)
			return true, values
		}
	}
	return false, values

}

type valueNode struct {
	depth int
	value int
}

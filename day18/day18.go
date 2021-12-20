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

func (*d) Run() (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var previousLine []*valueNode

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		values := parseLine(line)

		if previousLine != nil {
			values = add(previousLine, values)
			values = reduce(values)
		}
		previousLine = values

	}
	return magnitude(previousLine), 0
}

func magnitude(values []*valueNode) int {
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

func add(a, b []*valueNode) []*valueNode {
	for i := range a {
		a[i].depth++
	}
	for i := range b {
		b[i].depth++
	}

	return append(a, b...)
}

func parseLine(line string) (values []*valueNode) {
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
			values = append(values, &valueNode{depth, v})
			i += valueUntil - 1

		}
	}
	return
}

func reduce(values []*valueNode) []*valueNode {
	for {
		//fmt.Println("Values length:", len(values))
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
	}
}

func trySplit(values []*valueNode) (bool, []*valueNode) {
	for i := range values {
		if values[i].value > 9 {
			left := values[i].value / 2
			right := values[i].value - left

			values[i].value = left
			values[i].depth++
			newNode := &valueNode{values[i].depth, right}
			arrayEnd := append(
				[]*valueNode{newNode},
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

func tryExplode(values []*valueNode) (bool, []*valueNode) {
	for i := range values {
		if values[i].depth > 4 {
			goLeft := values[i].value
			if i > 0 {
				//fmt.Println("Left", i, ": Adding", goLeft, "to", values[i-1].value)
				values[i-1].value += goLeft
			}

			goRight := values[i+1].value
			if i < len(values)-2 {
				//fmt.Println("Right", i, ": Adding", goRight, "to", values[i+2].value)
				values[i+2].value += goRight

			}

			//fmt.Println("Setting", values[i+1].value, "to zero, depth", values[i+1].depth, "to", values[i+1].depth-1)
			values[i+1].depth--
			values[i+1].value = 0

			//fmt.Println("removing", values[i].value, "to from the array")

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

package day14

import (
	"bufio"
	"github.com/skamoen/advent2021/util"
	"log"
	"math"
	"os"
	"strings"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() (int, int) {
	//file, err := os.Open("./day14/example.txt")
	file, err := os.Open("./day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()

	var insert = make(map[string]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")
		insert[split[0]] = split[1]
	}

	newPolymer := polymer
	for n := 0; n < 10; n++ {
		for i := 0; i < len(polymer)-1; i++ {
			pair := polymer[i : i+2]
			lookup := insert[pair]
			newPolymer = newPolymer[:i+i+1] + lookup + newPolymer[i+i+1:]
		}
		polymer = newPolymer
	}

	min, max := minMax(countCharacters(polymer))

	return max - min, 0
}

func countCharacters(s string) map[string]int {
	var counter = make(map[string]int, 1)

	for i := range s {
		counter[string(s[i])]++
	}

	return counter
}

func minMax(m map[string]int) (int, int) {
	min := math.MaxInt
	max := 0
	for s := range m {
		if m[s] < min {
			min = m[s]
		}
		if m[s] > max {
			max = m[s]
		}
	}

	return min, max
}

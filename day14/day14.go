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

	var pairs = make(map[string]*pair, 1)
	var counter = make(map[string]int, 1)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")
		pairs[split[0]] = &pair{
			letters: [2]string{string(split[0][0]), string(split[0][1])},
			insert:  split[1],
			count:   0,
		}
	}

	for i := 0; i < len(polymer); i++ {
		if i < len(polymer)-1 {
			currentPair := polymer[i : i+2]
			pairs[currentPair].count++
		}
		counter[string(polymer[i])]++
	}

	partOne := 0
	for n := 1; n <= 40; n++ {
		var currentKeys = make([]string, 0, len(pairs))
		for s := range pairs {
			if pairs[s].count > 0 {
				currentKeys = append(currentKeys, s)
			}
		}

		for _, k := range currentKeys {
			currentPair := pairs[k]

			counter[currentPair.insert] += currentPair.count

			newPairA := currentPair.letters[0] + currentPair.insert
			pairs[newPairA].add += currentPair.count

			newPairB := currentPair.insert + currentPair.letters[1]
			pairs[newPairB].add += currentPair.count

			currentPair.count = 0
		}
		for _, p := range pairs {
			p.count += p.add
			p.add = 0
		}
		if n == 10 {
			min, max := minMax(counter)
			partOne = max - min
		}
	}

	min, max := minMax(counter)
	return partOne, max - min
}

type pair struct {
	letters [2]string
	insert  string
	count   int
	add     int
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

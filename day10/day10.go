package day10

import (
	"bufio"
	"github.com/skamoen/advent2021/util"
	"log"
	"os"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() (int, int) {
	file, err := os.Open("./day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	expected := map[string]string{
		"(": ")",
		"<": ">",
		"[": "]",
		"{": "}",
	}

	corruptScore := map[string]int{
		")": 3,
		">": 25137,
		"]": 57,
		"}": 1197,
	}

	completeScoreMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	opening := []string{"(", "{", "[", "<"}
	syntaxScore := 0
	var completeScores []int

lineScan:
	for scanner.Scan() {
		line := scanner.Text()

		var record []string
		for i := range line {
			currentCharacter := string(line[i])
			if util.ArrayContains(opening, currentCharacter) {
				record = append([]string{currentCharacter}, record...)
			} else if expected[record[0]] != currentCharacter {
				syntaxScore += corruptScore[currentCharacter]
				continue lineScan
			} else {
				record = record[1:]
			}
		}

		completeScore := 0
		for _, r := range record {
			completeScore = completeScore * 5
			completeScore += completeScoreMap[expected[r]]
		}
		completeScores = append(completeScores, completeScore)
	}

	return syntaxScore, util.Median(completeScores...)
}

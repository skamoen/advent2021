package day13

import (
	"bufio"
	"fmt"
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
	//file, err := os.Open("./day13/example.txt")
	file, err := os.Open("./day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var dots = make([][]bool, 1311)
	for i := range dots {
		dots[i] = make([]bool, 1311)
	}

	var folds []fold

	parseInstructions := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parseInstructions = true
			continue
		}

		if parseInstructions {
			split1 := strings.Split(line, " along ")
			split2 := strings.Split(split1[1], "=")
			axis, _ := strconv.Atoi(split2[1])
			folds = append(folds, fold{
				direction: split2[0],
				axis:      axis,
			})
			continue
		}

		path := strings.Split(line, ",")
		a, b := path[0], path[1]
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)

		dots[x][y] = true
	}

	nVisible := 0
	partOne := 0
	for i, f := range folds {
		nVisible = 0
		if f.direction == "y" {
			// fold up
			var newDots [][]bool = make([][]bool, len(dots))
			for i := range newDots {
				newDots[i] = make([]bool, f.axis)
			}

			for c := range dots {
				for r := range dots[c] {
					rBelow := f.axis + 1 + r
					rOrig := f.axis - 1 - r
					belowFold := dots[c][rBelow]
					foldOn := dots[c][rOrig]
					newDots[c][rOrig] = belowFold || foldOn

					if belowFold || foldOn {
						nVisible++
					}
					if rOrig == 0 {
						break
					}
				}
			}
			dots = newDots
		} else if f.direction == "x" {
			var newDots = make([][]bool, f.axis)
			for i := range newDots {
				newDots[i] = make([]bool, len(dots[i]))
			}

			for c := range dots[:f.axis] {
				for r := range dots[c] {
					cRight := f.axis + 1 + c
					cOrig := f.axis - 1 - c
					rightFold := dots[cRight][r]
					foldOn := dots[cOrig][r]
					newDots[cOrig][r] = rightFold || foldOn

					if rightFold || foldOn {
						nVisible++
					}
				}
			}
			dots = newDots
		}
		if i == 0 {
			partOne = nVisible
		}
	}

	for j := 0; j < len(dots[0]); j++ {
		for i := 0; i < len(dots); i++ {
			if dots[i][j] {
				fmt.Print("X ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	return partOne, 0
}

type fold struct {
	direction string
	axis      int
}

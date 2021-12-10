package day04

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
	file, err := os.Open("./day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	drawings := scanner.Text()
	scanner.Scan()

	var boards = []board{{columnCounter: make([]int, 5), rowCounter: make([]int, 5)}}
	currentBoard := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currentBoard++
			boards = append(boards, board{columnCounter: make([]int, 5), rowCounter: make([]int, 5)})
			continue
		}

		nums := strings.Fields(line)
		boards[currentBoard].rows = append(boards[currentBoard].rows, nums)

	}

	boardsWon := make([]bool, len(boards))
	var boardScores []int
	for _, drawing := range strings.Split(drawings, ",") {
		for i := range boards {
			for j := range boards[i].rows {
				for k := range boards[i].rows[j] {
					// Mark number
					if boards[i].rows[j][k] == drawing {
						boards[i].rows[j][k] = "-1"
						boards[i].rowCounter[j]++
						boards[i].columnCounter[k]++

						if boards[i].rowCounter[j] == 5 || boards[i].columnCounter[k] == 5 {
							if !boardsWon[i] {
								// Sum the score
								boardSum := 0
								for x := range boards[i].rows {
									for y := range boards[i].rows[x] {
										parseInt, _ := strconv.Atoi(boards[i].rows[x][y])
										if parseInt > 0 {
											boardSum += parseInt
										}
									}
								}
								parseDraw, _ := strconv.Atoi(drawing)
								boardsWon[i] = true
								boardScores = append(boardScores, boardSum*parseDraw)
							}
						}
					}
				}
			}
		}
	}
	return boardScores[0], boardScores[len(boardScores)-1]
}

type board struct {
	rows          [][]string
	rowCounter    []int
	columnCounter []int
}

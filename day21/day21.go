package day21

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

var inputFile = "./day21/input.txt"

var states map[int][2]int
var rolls map[int]int

func (*d) Run() (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 0 = player 1 starting
	// 1 = player 2 starting
	var gamedata []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		startingPosition, _ := strconv.Atoi(string(line[len(line)-1]))
		gamedata = append(gamedata, startingPosition)
	}

	gamedata = append(gamedata, 0, 0)

	states = make(map[int][2]int)
	rolls = make(map[int]int, 27)
	preroll()
	p1Wins, p2Wins := roll(gamedata[0], gamedata[1], 0, 0)

	var partTwo = 0
	if p1Wins > p2Wins {
		partTwo = p1Wins
	} else {
		partTwo = p2Wins
	}

	diceRolls := 0
	currentPlayer := 0 // Gamedata index

	for i := 1; gamedata[2] < 1000 && gamedata[3] < 1000; i = i + 3 {
		roll := 3*i + 3

		// Position
		newPosition := gamedata[currentPlayer] + roll%10
		if newPosition > 10 {
			newPosition = newPosition % 10
		}
		gamedata[currentPlayer] = newPosition
		// score
		gamedata[currentPlayer+2] += gamedata[currentPlayer]

		if currentPlayer == 0 {
			currentPlayer = 1
		} else {
			currentPlayer = 0
		}
		diceRolls += 3
	}

	partOne := 0
	if currentPlayer == 1 {
		partOne = gamedata[3] * diceRolls
	} else {
		partOne = gamedata[2] * diceRolls
	}

	return partOne, partTwo
}

func preroll() {
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				rolls[i+j+k]++
			}
		}
	}
}
func roll(onePosition, twoPosition, oneScore, twoScore int) (int, int) {
	if oneScore >= 21 {
		return 1, 0
	} else if twoScore >= 21 {
		return 0, 1
	}
	// "Hash" of the gamestate
	state := onePosition<<32 +
		twoPosition<<16 +
		oneScore<<8 +
		twoScore

	if outcome, ok := states[state]; ok {
		return outcome[0], outcome[1]
	}

	var p1Wins, p2Wins int
	for sum, freq := range rolls {
		// Position
		newPosition := onePosition + sum
		if newPosition > 10 {
			newPosition = newPosition % 10
		}

		p2Win, p1Win := roll(twoPosition, newPosition, twoScore, oneScore+newPosition)
		p1Wins += p1Win * freq
		p2Wins += p2Win * freq
	}
	states[state] = [2]int{p1Wins, p2Wins}
	return p1Wins, p2Wins
}

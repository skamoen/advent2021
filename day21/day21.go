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

//var inputFile = "./day21/example.txt"

var scores map[int]int

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

	scores = make(map[int]int, 2)
	a := player{
		position: 4,
	}
	b := player{
		position: 8,
	}
	roll(a, b, 1, 1)
	roll(a, b, 2, 1)
	roll(a, b, 3, 1)

	return partOne, 0
}
func roll(one, two player, thisRoll int, turn int) {
	var currentPlayer *player
	if turn%2 != 0 {
		currentPlayer = &one
	} else {
		currentPlayer = &two
	}

	// Position
	newPosition := currentPlayer.position + thisRoll
	if newPosition > 10 {
		newPosition = newPosition % 10
	}
	currentPlayer.score += newPosition

	if currentPlayer.score >= 21 {
		scores[turn%2]++
		return
	} else {
		roll(one, two, 1, turn+1)
		roll(one, two, 2, turn+1)
		roll(one, two, 3, turn+1)
	}
}

type player struct {
	score    int
	position int
}

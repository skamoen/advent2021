package day02

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

func (*d) Run() {
	file, err := os.Open("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var commands [][]string
	var depth, horizontalPosition int64 = 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		commands = append(commands, command)
		movement := command[0]
		amount, _ := strconv.ParseInt(command[1], 10, 64)
		//log.Println("Going", movement, "amount", amount)
		if movement == "forward" {
			horizontalPosition = horizontalPosition + amount
		}
		if movement == "down" {
			depth = depth + amount
		}
		if movement == "up" {
			depth = depth - amount
		}

	}

	log.Println("------- PART ONE ---------")
	log.Println("Command length", len(commands))
	log.Println("Horizontal", horizontalPosition, "Depth", depth, "Answer", horizontalPosition*depth)

	var hPos2, depth2, aim int64 = 0, 0, 0
	for _, c := range commands {
		movement := c[0]
		amount, _ := strconv.ParseInt(c[1], 10, 64)
		if movement == "down" {
			aim = aim + amount
		}
		if movement == "up" {
			aim = aim - amount
		}
		if movement == "forward" {
			hPos2 = hPos2 + amount
			depth2 = depth2 + (aim * amount)
		}
	}
	log.Println("---- PART TWO ----")
	log.Println("Horizontal 2", hPos2, "Depth", depth2, "Answer", hPos2*depth2)

}

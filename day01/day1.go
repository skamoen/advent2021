package day01

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

func (*d) Run() (int, int) {
	file, err := os.Open("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	deeper := 0
	var measurements []int

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	prevReading, _ := strconv.Atoi(scanner.Text())
	measurements = append(measurements, prevReading)
	for scanner.Scan() {
		reading, _ := strconv.Atoi(scanner.Text())
		if reading > prevReading {
			deeper++
		}

		measurements = append(measurements, reading)
		prevReading = reading
	}

	deeperWindows := 0
	var prevSum int
	for i := 0; i <= len(measurements)-3; i++ {
		currentWindow := measurements[i : i+3]
		currentSum := 0
		for _, m := range currentWindow {
			currentSum = currentSum + m
		}
		if prevSum != 0 {
			if currentSum > prevSum {
				deeperWindows++
			}
		}
		prevSum = currentSum
	}

	return deeper, deeperWindows
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	day1()
}

func day1() {
	file, err := os.Open("./1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	deeper := 0
	var measurements []int64

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	prevReading, err := strconv.ParseInt(scanner.Text(), 10, 64)
	measurements = append(measurements, prevReading)
	for scanner.Scan() {
		reading, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if reading > prevReading {
			deeper++
		}

		measurements = append(measurements, reading)
		prevReading = reading
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	deeperWindows := 0
	var prevSum int64
	for i := 0; i <= len(measurements)-3; i++ {
		currentWindow := measurements[i : i+3]
		var currentSum int64 = 0
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

	fmt.Println("Readings deeper: ", deeper)
	fmt.Println("Readings deeper windows: ", deeperWindows)
}

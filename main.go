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

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	prevReading, err := strconv.ParseInt(scanner.Text(), 10, 64)
	for scanner.Scan() {
		reading, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if reading > prevReading {
			deeper++
		}
		prevReading = reading
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Readings deeper: ", deeper)
}

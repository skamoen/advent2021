package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//day1()
	//day2()
	day3()
}

func day3() {
	file, err := os.Open("./3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bits [][]string = make([][]string, 12)
	var part2Bits [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		binary := scanner.Text()
		var row []string
		for i := range binary {
			bits[i] = append(bits[i], string(binary[i]))
			row = append(row, string(binary[i]))
		}
		part2Bits = append(part2Bits, row)

	}

	var gamma string
	var epsilon string

	for _, bitArray := range bits {
		nZero, nOne := 0, 0
		for _, b := range bitArray {
			if b == "1" {
				nOne++
			} else if b == "0" {
				nZero++
			}
		}

		if nZero > nOne {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaValue, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonValue, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Part one result", gammaValue*epsilonValue)

	oxygen := filterOxygen(part2Bits, 0)
	co2 := filterCo2(part2Bits, 0)

	oxyString := ""
	for _, o := range oxygen[0] {
		oxyString += o
	}
	co2String := ""
	for _, o := range co2[0] {
		co2String += o
	}

	oxyValue, _ := strconv.ParseInt(oxyString, 2, 64)
	co2Value, _ := strconv.ParseInt(co2String, 2, 64)

	log.Println("result", oxyValue*co2Value)

}

func day2() {
	file, err := os.Open("./2/input.txt")
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

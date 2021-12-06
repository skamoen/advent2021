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
	//day3()
	//day4()
	//day5()
	day6()
}

func day6() {
	file, err := os.Open("./6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fish []int

	fishPerDay := make([]int, 256)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, ",")
		for _, s := range f {
			i, _ := strconv.Atoi(s)
			fish = append(fish, i)
			fishPerDay[i]++
		}
	}

	for i := 0; i < 80; i++ {
		nFish := len(fish)
		for i := 0; i < nFish; i++ {
			if fish[i] == 0 {
				fish[i] = 6
				fish = append(fish, 8)
				continue
			}
			fish[i]--
		}
	}
	log.Println("No. fish after 80  days", len(fish))

	for i := 0; i < 256; i++ {
		fishOnZero := fishPerDay[0]
		newFishPerDay := append(fishPerDay[1:], 0)
		newFishPerDay[6] += fishOnZero
		newFishPerDay[8] = fishOnZero
		fishPerDay = newFishPerDay
	}

	nFish := 0
	for _, n := range fishPerDay {
		nFish += n
	}
	log.Println("No. Fish after 256 days", nFish)
}

func day5() {
	file, err := os.Open("./5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []vector
	var diag []vector

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")
		from, to := split[0], split[1]
		v := makeVectorFromString(from, to)
		if string(v.direction[0]) != "d" {
			lines = append(lines, v)
		} else {
			diag = append(diag, v)
		}
	}

	var m [1000][1000]int
	for _, line := range lines {

		x, y := line.start[0], line.start[1]
		for j := 0; j <= line.length; j++ {
			if line.direction == "h" {
				m[x+j][y]++
			} else if line.direction == "v" {
				m[x][y+j]++
			}
		}
	}

	overlap := 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] >= 2 {
				overlap++
			}
		}
	}

	log.Println("Part 1 Overlapping", overlap)

	for _, line := range diag {
		x, y := line.start[0], line.start[1]
		for j := 0; j <= line.length; j++ {
			if line.direction == "du" {
				m[x+j][y+j]++
			} else {
				m[x+j][y-j]++
			}
		}
	}

	overlap = 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] >= 2 {
				overlap++
			}
		}
	}

	log.Println("Part 2 Overlapping", overlap)
}

type board struct {
	rows          [][]string
	rowCounter    []int
	columnCounter []int
}

func day4() {
	file, err := os.Open("./4/input.txt")
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
								log.Println("Result for board", i, ":", boardSum*parseDraw)
								boardsWon[i] = true
							}
						}
					}
				}
			}
		}
	}
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

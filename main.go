package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	//day1()
	//day2()
	//day3()
	//day4()
	//day5()
	//day6()
	//day7()
	//day8()
	//day9()
	day10()

	diff := time.Now().Sub(start)
	fmt.Println("Took", diff.Microseconds(), "microseconds")
}

func day10() {
	file, err := os.Open("./10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	expected := map[string]string{
		"(": ")",
		"<": ">",
		"[": "]",
		"{": "}",
	}

	corruptScore := map[string]int{
		")": 3,
		">": 25137,
		"]": 57,
		"}": 1197,
	}

	completeScoreMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	opening := []string{"(", "{", "[", "<"}
	syntaxScore := 0
	var completeScores []int

lineScan:
	for scanner.Scan() {
		line := scanner.Text()

		var record []string
		for i := range line {
			currentCharacter := string(line[i])
			if arrayContains(opening, currentCharacter) {
				record = append([]string{currentCharacter}, record...)
			} else if expected[record[0]] != currentCharacter {
				syntaxScore += corruptScore[currentCharacter]
				continue lineScan
			} else {
				record = record[1:]
			}
		}

		completeScore := 0
		for _, r := range record {
			completeScore = completeScore * 5
			completeScore += completeScoreMap[expected[r]]
		}
		completeScores = append(completeScores, completeScore)
	}

	fmt.Println("Part one corruptScore", syntaxScore, "Autocomplete score", median(completeScores...))
}

func day9() {
	file, err := os.Open("./9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var heights [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]int, len(line))
		for i := range line {
			c, _ := strconv.Atoi(string(line[i]))
			row[i] = c
		}
		heights = append(heights, row)
	}

	riskSum := 0
	var sizes []int
	for i := range heights {
		for j := range heights[i] {
			currentHeight := heights[i][j]

			// Check up
			if j > 0 {
				if heights[i][j-1] == currentHeight || minInt(currentHeight, heights[i][j-1]) != currentHeight {
					continue
				}
			}
			// Check right
			if i < len(heights)-1 {
				if heights[i+1][j] == currentHeight || minInt(currentHeight, heights[i+1][j]) != currentHeight {
					continue
				}
			}
			// Check bottom
			if j < len(heights[i])-1 {
				if heights[i][j+1] == currentHeight || minInt(currentHeight, heights[i][j+1]) != currentHeight {
					continue
				}
			}
			// Check left
			if i > 0 {
				if heights[i-1][j] == currentHeight || minInt(currentHeight, heights[i-1][j]) != currentHeight {
					continue
				}
			}
			riskSum += currentHeight + 1
			b := &basin{start: []int{i, j}}
			size := b.findSize(heights)
			sizes = append(sizes, size)
		}
	}

	sort.Ints(sizes)
	ints := sizes[len(sizes)-3:]
	fmt.Println("Riskfactor", riskSum, "Basin size", ints[0]*ints[1]*ints[2])
}

func (b *basin) findSize(heights [][]int) int {
	b.marker = make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		b.marker[i] = make([]bool, len(heights))
	}

	// Check
	b.checkPoint(b.start[0], b.start[1], heights)

	return b.size
}

func (b *basin) checkPoint(i, j int, heights [][]int) {

	if b.marker[i][j] == false {
		b.size++
		b.marker[i][j] = true
	} else {
		return
	}

	// Check left
	if i > 0 {
		if heights[i-1][j] != 9 {
			b.checkPoint(i-1, j, heights)
		}
	}
	// Check up
	if j > 0 {
		if heights[i][j-1] != 9 {
			b.checkPoint(i, j-1, heights)
		}
	}
	// Check right
	if i < len(heights)-1 {
		if heights[i+1][j] != 9 {
			b.checkPoint(i+1, j, heights)
		}
	}
	// Check bottom
	if j < len(heights[i])-1 {
		if heights[i][j+1] != 9 {
			b.checkPoint(i, j+1, heights)
		}
	}
}

type basin struct {
	start  []int
	marker [][]bool // Mark X, Y as included
	size   int
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func day8() {
	start := time.Now()
	file, err := os.Open("./8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nKnown := 0
	totalValue := 0
	for scanner.Scan() {
		line := scanner.Text()
		d := &display{knownSignals: make(map[int][]string, 10), solved: make(map[int]bool, 10)}

		f := strings.Split(line, " | ")

		for _, digit := range strings.Fields(f[0]) {
			da := stringToCharArray(digit)
			d.wires = append(d.wires, da)
			d.setUniqeMappings(da)
		}

		for _, o := range strings.Fields(f[1]) {
			d.output = append(d.output, stringToCharArray(o))
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				nKnown++
			}
		}
		d.solve()
		totalValue += d.value()
	}

	fmt.Println("Part one: ", nKnown, "Part two: ", totalValue)

	diff := time.Now().Sub(start)
	fmt.Println("Took", diff.Microseconds(), "microseconds")

}

func day7() {
	file, err := os.Open("./7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var crabs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, ",")
		for _, s := range f {
			i, _ := strconv.Atoi(s)
			crabs = append(crabs, i)
		}
	}

	median := median(crabs...)

	fuel1, fuel2 := 0, 0
	for _, c := range crabs {
		fuel1 += abs(c - median)
		n := abs(c - avg(crabs...))
		fuel2 += (n + 1) * n / 2
	}
	fmt.Println("Fuel spent: ", fuel1, fuel2)
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	} else {
		return i
	}
}

func median(n ...int) int {
	sort.Ints(n)
	m := len(n) / 2
	if !(len(n)%2 == 0) {
		return n[m]
	}
	return (n[m-1] + n[m]) / 2
}

func avg(n ...int) int {
	sum := sumArray(n)
	l := len(n)

	avg := sum / l
	return avg
}

func day6() {

	start := time.Now()

	file, err := os.Open("./6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fishPerDay := make([]int, 256+9)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, ",")
		for _, s := range f {
			i, _ := strconv.Atoi(s)
			fishPerDay[i]++
		}
	}

	day80 := 0
	for i := 0; i < 256; i++ {
		fishPerDay[i+7] += fishPerDay[i]
		fishPerDay[i+9] = fishPerDay[i]
		if i == 79 {
			day80 = sumArray(fishPerDay[80:])
		}
	}
	fmt.Println("No. Fish after 80 & 256 days", day80, sumArray(fishPerDay[256:]))
	took := time.Now().Sub(start)
	fmt.Println(took.Microseconds())
}

func sumArray(a []int) int {
	s := 0
	for i := range a {
		s += a[i]
	}
	return s
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

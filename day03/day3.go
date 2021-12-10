package day03

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

func (*d) Run() {
	file, err := os.Open("./day03/input.txt")
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

func filterOxygen(bits [][]string, pos int) [][]string {
	common := util.MostCommon(bits, pos)

	var remainingbits [][]string
	for i := 0; i < len(bits); i++ {
		if bits[i][pos] == common {
			remainingbits = append(remainingbits, bits[i])
		}
	}
	if len(remainingbits) == 1 {
		return remainingbits
	}

	return filterOxygen(remainingbits, pos+1)
}

func filterCo2(bits [][]string, pos int) [][]string {

	common := util.LeastCommon(bits, pos)

	var remainingbits [][]string
	for i := 0; i < len(bits); i++ {
		if bits[i][pos] == common {
			remainingbits = append(remainingbits, bits[i])
		}
	}
	if len(remainingbits) == 1 {
		return remainingbits
	}

	return filterCo2(remainingbits, pos+1)
}

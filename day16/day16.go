package day16

import (
	"bufio"
	"github.com/skamoen/advent2021/util"
	"log"
	"math"
	"os"
	"strconv"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

var versions int64 = 0
var inputFile = "./day16/input.txt"

func (*d) Run() (int, int) {
	//file, err := os.Open("./day16/example.txt")
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hexConvert := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	versions = 0
	binaryString := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rawData := make([]byte, len(line)*4)
		for i := range line {
			u := string(line[i])
			copy(rawData[i*4:i*4+4], hexConvert[u])
		}
		binaryString = string(rawData)
	}

	_, p := parsePacket(binaryString, 0)

	return int(versions), int(p.Value)
}

func parsePacket(binaryString string, i int) (int, *packet) {
	p := &packet{}
	packetVersion, _ := strconv.ParseInt(binaryString[i:i+3], 2, 64)
	versions += packetVersion
	i = i + 3
	packetTypeId, _ := strconv.ParseInt(binaryString[i:i+3], 2, 64)
	i = i + 3

	p.Version = packetVersion
	p.TypeId = packetTypeId

	if packetTypeId == 4 {
		// Parse groups of 4 bits
		prefix := ""
		value := ""
		for j := 0; prefix != "0"; j++ {
			prefix = string(binaryString[i])
			i++
			value += binaryString[i : i+4]
			i = i + 4
		}
		packetValue, _ := strconv.ParseInt(value, 2, 64)
		p.Value = packetValue
		return i, p
	} else {
		lengthTypeId := string(binaryString[i])
		i++
		if lengthTypeId == "0" {
			nBitsSubPackets, _ := strconv.ParseInt(binaryString[i:i+15], 2, 64)
			i = i + 15
			for j := 0; j < int(nBitsSubPackets); {
				newI, subPacket := parsePacket(binaryString, i)
				j += newI - i
				i = newI
				p.SubPackets = append(p.SubPackets, subPacket)
			}
		} else if lengthTypeId == "1" {
			nSubPackets, _ := strconv.ParseInt(binaryString[i:i+11], 2, 64)
			i = i + 11
			for k := 0; k < int(nSubPackets); k++ {
				newI, subPacket := parsePacket(binaryString, i)
				i = newI
				p.SubPackets = append(p.SubPackets, subPacket)
			}
		}

		switch packetTypeId {
		case 0:
			var sum int64 = 0
			for s := range p.SubPackets {
				sum += p.SubPackets[s].Value
			}
			p.Value = sum
		case 1:
			var product int64 = 1
			for s := range p.SubPackets {
				product = product * p.SubPackets[s].Value
			}
			p.Value = product
		case 2:
			var min int64 = math.MaxInt
			for s := range p.SubPackets {
				if p.SubPackets[s].Value < min {
					min = p.SubPackets[s].Value
				}
			}
			p.Value = min
		case 3:
			var max int64 = 0
			for s := range p.SubPackets {
				if p.SubPackets[s].Value > max {
					max = p.SubPackets[s].Value
				}
			}
			p.Value = max
		case 5:
			if p.SubPackets[0].Value > p.SubPackets[1].Value {
				p.Value = 1
			} else {
				p.Value = 0
			}
		case 6:
			if p.SubPackets[0].Value < p.SubPackets[1].Value {
				p.Value = 1
			} else {
				p.Value = 0
			}
		case 7:
			if p.SubPackets[0].Value == p.SubPackets[1].Value {
				p.Value = 1
			} else {
				p.Value = 0
			}
		}
	}
	return i, p
}

type packet struct {
	Version    int64
	TypeId     int64
	Value      int64
	SubPackets []*packet
}

package day20

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

var inputFile = "./day20/input.txt"

//var inputFile = "./day20/example.txt"

type imageType [][]int

var iea string

func (*d) Run() (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var image = make(imageType, 0, 128)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	iea = scanner.Text() // Skip newline
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for c := range line {
			if string(line[c]) == "." {
				row[c] = 0
			} else if string(line[c]) == "#" {
				row[c] = 1
			}
		}
		image = append(image, row)
	}

	var partOne = 0
	var partTwo = 0
	for i := 0; i < 50; i++ {
		var lightPixels = 0
		image = image.expandCanvas()

		var newImage = make(imageType, len(image))
		for r := 0; r < len(image); r++ {
			row := make([]int, len(image[r]))
			for c := 0; c < len(image[r]); c++ {
				enhance := image.enhance(r, c, image[0][0])
				if enhance == 1 {
					lightPixels++
				}
				row[c] = enhance
			}
			newImage[r] = row
		}
		if i == 1 {
			partOne = lightPixels
		} else if i == 49 {
			partTwo = lightPixels
		}
		image = newImage
	}

	return partOne, partTwo
}

func (i imageType) expandCanvas() imageType {
	defaultValue := i[0][0]
	defaultRow := make([]int, len(i[0])+2)
	if defaultValue != 0 {
		for j := range defaultRow {
			defaultRow[j] = defaultValue
		}
	}

	newImage := make(imageType, len(i)+2)
	newImage[0] = make([]int, len(i[0])+2)
	newImage[len(newImage)-1] = make([]int, len(i[0])+2)
	if defaultValue != 0 {
		copy(newImage[0], defaultRow)
		copy(newImage[len(newImage)-1], defaultRow)
	}

	for r := range i {
		newRow := make([]int, len(i[r])+2)
		newRow[0] = defaultValue
		copy(newRow[1:], i[r])
		newRow[len(newRow)-1] = defaultValue
		newImage[r+1] = newRow
	}

	return newImage
}

func (i imageType) enhance(r, c int, d int) int {
	var neighbours = [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}
	var lookup string

	for _, neighbour := range neighbours {
		if r+neighbour[0] < 0 || r+neighbour[0] >= len(i) || c+neighbour[1] < 0 || c+neighbour[1] >= len(i[r]) {
			lookup += strconv.Itoa(d)
		} else {
			n := i[r+neighbour[0]][c+neighbour[1]]
			lookup += strconv.Itoa(n)
		}
	}

	index, err := strconv.ParseInt(lookup, 2, 32)
	if err != nil {
		log.Fatal("Cant convert lookup")
	}

	if string(iea[index]) == "#" {
		return 1
	} else {
		return 0
	}
}

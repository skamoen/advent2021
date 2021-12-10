package day06

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

func (*d) Run() (int, int) {
	file, err := os.Open("./day06/input.txt")
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
			day80 = util.SumArrayInts(fishPerDay[80:])
		}
	}
	return day80, util.SumArrayInts(fishPerDay[256:])
}

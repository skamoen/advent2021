package day06

import (
	"bufio"
	"fmt"
	"github.com/skamoen/advent2021/util"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() {

	start := time.Now()

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
	fmt.Println("No. Fish after 80 & 256 days", day80, util.SumArrayInts(fishPerDay[256:]))
	took := time.Now().Sub(start)
	fmt.Println(took.Microseconds())
}

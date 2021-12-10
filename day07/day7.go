package day07

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
	file, err := os.Open("./day07/input.txt")
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

	median := util.Median(crabs...)

	fuel1, fuel2 := 0, 0
	for _, c := range crabs {
		fuel1 += util.Abs(c - median)
		n := util.Abs(c - util.Avg(crabs...))
		fuel2 += (n + 1) * n / 2
	}
	return fuel1, fuel2
}

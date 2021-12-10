package main

import (
	"fmt"
	"github.com/skamoen/advent2021/day01"
	"github.com/skamoen/advent2021/day02"
	"github.com/skamoen/advent2021/day03"
	"github.com/skamoen/advent2021/day04"
	"github.com/skamoen/advent2021/day05"
	"github.com/skamoen/advent2021/day06"
	"github.com/skamoen/advent2021/day07"
	"github.com/skamoen/advent2021/day08"
	"github.com/skamoen/advent2021/day09"
	"github.com/skamoen/advent2021/day10"
	"github.com/skamoen/advent2021/util"
	"time"
)

func main() {
	var days = []util.Entry{util.Nop(),
		day01.Get(),
		day02.Get(),
		day03.Get(),
		day04.Get(),
		day05.Get(),
		day06.Get(),
		day07.Get(),
		day08.Get(),
		day09.Get(),
		day10.Get(),
	}

	fmt.Println("--- TODAY ---")
	start := time.Now()
	if len(days) <= start.Day() {
		fmt.Println("Day", start.Day(), "not implemented yet")
		return
	}
	part1, part2 := days[start.Day()].Run()
	fmt.Println("Day", start.Day(), "\tPart 1", part1, "\tPart 2", part2)
	diff := time.Now().Sub(start)
	fmt.Println("Took", diff.Microseconds(), "microseconds")

	if true {
		fmt.Println("\n--- BENCHMARK ---")

		startTotal := time.Now()
		for i, f := range days[1:] {
			part1, part2 := f.Run()
			fmt.Println("Day", i+1, "\tPart 1", part1, "\tPart 2", part2)
		}
		diffTotal := time.Now().Sub(startTotal)
		fmt.Println("Took Total", diffTotal.Milliseconds(), "milliseconds")
	}
}

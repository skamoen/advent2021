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
	"github.com/skamoen/advent2021/day11"
	"github.com/skamoen/advent2021/day12"
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
		day11.Get(),
		day12.Get(),
	}

	fmt.Println("--- TODAY ---")
	start := time.Now()
	today := start.Day()
	if len(days) <= today {
		fmt.Println("Day", today, "not implemented yet")
		return
	}
	part1, part2 := days[today].Run()
	fmt.Println("Day", today, "\tPart 1", part1, "\tPart 2", part2)
	diff := time.Now().Sub(start)
	fmt.Println("Took", diff.Microseconds(), "microseconds")

	if false {
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

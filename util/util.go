package util

import (
	"fmt"
	"sort"
)

var binarySum []int = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

type Entry interface {
	Run() (int, int)
}

type d struct {
}

func Nop() Entry {
	return &d{}
}

func (*d) Run() (int, int) {
	return 0, 0
}

func PrintGrid(g [][]int) {
	for i := range g {
		for j := range g[i] {
			fmt.Print(g[i][j], " ")
		}
		fmt.Print("\n")
	}

}

func PrintGridImage(g [][]int) {
	for i := range g {
		for j := range g[i] {
			if g[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()

}

func UintBinaryToDecimal(b []uint8) int {
	sum := 0
	for i := range b {
		sum += int(b[len(b)-1-i]) * binarySum[i]
	}
	return sum
}

func MostCommon(bits [][]string, pos int) string {
	zeroes, ones := Count(bits, pos)
	if zeroes == ones {
		return "1"
	} else if zeroes > ones {
		return "0"
	} else {
		return "1"
	}
}

func LeastCommon(bits [][]string, pos int) string {
	zeroes, ones := Count(bits, pos)
	if zeroes > ones {
		return "1"
	} else {
		return "0"
	}
}

func Count(bits [][]string, pos int) (int, int) {
	nZero, nOne := 0, 0
	for i := 0; i < len(bits); i++ {
		if bits[i][pos] == "1" {
			nOne++
		} else if bits[i][pos] == "0" {
			nZero++
		}
	}
	return nZero, nOne
}

func SumArrayInts(a []int) int {
	s := 0
	for i := range a {
		s += a[i]
	}
	return s
}

func Abs(i int) int {
	if i < 0 {
		return -1 * i
	} else {
		return i
	}
}

func Median(n ...int) int {
	sort.Ints(n)
	m := len(n) / 2
	if !(len(n)%2 == 0) {
		return n[m]
	}
	return (n[m-1] + n[m]) / 2
}

func Avg(n ...int) int {
	sum := SumArrayInts(n)
	l := len(n)

	avg := sum / l
	return avg
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func ArrayContainsAll(a []string, b []string) bool {
	for _, s := range b {
		if !ArrayContains(a, s) {
			return false
		}
	}
	return true
}

func ArrayContainsExact(a []string, b []string) bool {
	return len(a) == len(b) && ArrayContainsAll(a, b)
}

func StringToCharArray(s string) []string {
	a := make([]string, len(s))
	for i := range s {
		a[i] = string(s[i])
	}
	return a
}

func ArrayContains(s []string, str string) bool {
	for i, _ := range s {
		if s[i] == str {
			return true
		}
	}
	return false
}

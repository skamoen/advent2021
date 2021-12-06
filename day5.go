package main

import (
	"strconv"
	"strings"
)

func makeVectorFromString(from, to string) vector {
	fromSplit := strings.Split(from, ",")
	toSplit := strings.Split(to, ",")

	x1, _ := strconv.Atoi(fromSplit[0])
	y1, _ := strconv.Atoi(fromSplit[1])
	x2, _ := strconv.Atoi(toSplit[0])
	y2, _ := strconv.Atoi(toSplit[1])

	return makeVector(x1, y1, x2, y2)
}

func makeVector(x1, y1, x2, y2 int) vector {
	v := vector{}

	if x1 == x2 {
		v.direction = "v"
		if y1 < y2 {
			v.start = [2]int{x1, y1}
			v.length = y2 - y1
		} else {
			v.start = [2]int{x2, y2}
			v.length = y1 - y2
		}
	} else if y1 == y2 {
		v.direction = "h"
		if x1 < x2 {
			v.start = [2]int{x1, y1}
			v.length = x2 - x1
		} else {
			v.start = [2]int{x2, y2}
			v.length = x1 - x2
		}
	} else {
		if x1 < x2 {
			v.start = [2]int{x1, y1}
			v.length = x2 - x1
			if y1 < y2 {
				v.direction = "du"
			} else {
				v.direction = "dd"
			}
		} else {
			v.start = [2]int{x2, y2}
			v.length = x1 - x2
			if y1 < y2 {
				v.direction = "dd"
			} else {
				v.direction = "du"
			}
		}
	}
	return v
}

type vector struct {
	direction string
	start     [2]int
	length    int
}

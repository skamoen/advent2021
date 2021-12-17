package day17

import (
	"github.com/skamoen/advent2021/util"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

var (
	minX int
	maxX int
	minY int
	maxY int
	//inputFile = "./day17/input.txt"
)

// target area: x=281..311, y=-74..-54

func (*d) Run() (int, int) {
	// Actual input
	//target area: x=281..311, y=-74..-54
	minX, maxX, minY, maxY = 281, 311, -74, -54

	// Example
	//target area: x=20..30, y=-10..-5
	//minX, maxX, minY, maxY = 20, 30, -10, -5

	x := 0
	y := 0

	vX, vY := 6, 9
	highest := y
	nVelocity := 0

	for dx := 1; dx <= maxX; dx++ {
		for dy := minY; dy < maxX; dy++ {
			vX = dx
			vY = dy

			x, y = 0, 0

			currentTop := 0

			target, abort := false, false
			for i := 0; !target && !abort; i++ {
				// Do step
				x = x + vX
				y = y + vY

				// Drag
				if vX > 0 {
					vX--
				} else if vX < 0 {
					vX++
				}

				// Gravity
				vY--

				if y > currentTop {
					currentTop = y
				}
				target, abort = inTargetArea(x, y)

				if target {
					nVelocity++
					if currentTop > highest {
						highest = currentTop
					}
				}
			}
		}
	}

	return highest, nVelocity
}

func inTargetArea(x, y int) (bool, bool) {
	if x <= maxX && x >= minX && y <= maxY && y >= minY {
		return true, false
	}
	if x > maxX {
		return false, true
	}
	if y < minY {
		return false, true
	}

	return false, false
}

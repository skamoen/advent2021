package day08

import (
	"bufio"
	"github.com/skamoen/advent2021/util"
	"log"
	"math"
	"os"
	"strings"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() (int, int) {
	file, err := os.Open("./day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nKnown := 0
	totalValue := 0
	for scanner.Scan() {
		line := scanner.Text()
		d := &display{knownSignals: make(map[int][]string, 10), solved: make(map[int]bool, 10)}

		f := strings.Split(line, " | ")

		for _, digit := range strings.Fields(f[0]) {
			da := util.StringToCharArray(digit)
			d.wires = append(d.wires, da)
			d.setUniqeMappings(da)
		}

		for _, o := range strings.Fields(f[1]) {
			d.output = append(d.output, util.StringToCharArray(o))
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				nKnown++
			}
		}
		d.solve()
		totalValue += d.value()
	}

	return nKnown, totalValue
}

func (d *display) value() int {
	if len(d.knownSignals) != 10 {
		log.Println("Not solved yet")
		return 0
	}

	value := 0
	for i, o := range d.output {
		for v, s := range d.knownSignals {
			if util.ArrayContainsExact(o, s) {
				value += int(math.Pow10(3-i)) * v
			}
		}
	}
	return value
}

func (d *display) solve() {
	for i := 0; len(d.knownSignals) < 10; i++ {
		for i, wire := range d.wires {
			if d.solved[i] {
				continue
			}
			if len(wire) == 5 {
				if util.ArrayContainsAll(wire, d.knownSignals[1]) {
					// It's 3
					d.knownSignals[3] = wire
					d.solved[i] = true
				} else if _, ok := d.knownSignals[6]; ok {
					if util.ArrayContainsAll(d.knownSignals[6], wire) {
						// It's 5
						d.knownSignals[5] = wire
						d.solved[i] = true
					} else {
						// It's 2
						d.knownSignals[2] = wire
						d.solved[i] = true
					}
				}
			} else if len(wire) == 6 {
				// 6 or 9
				if _, ok := d.knownSignals[3]; ok {
					if util.ArrayContainsAll(wire, d.knownSignals[3]) {
						// It's 9
						d.knownSignals[9] = wire
						d.solved[i] = true
					} else {
						if _, ok := d.knownSignals[7]; ok {
							if util.ArrayContainsAll(wire, d.knownSignals[7]) {
								// It's 0
								d.knownSignals[0] = wire
								d.solved[i] = true
							} else {
								// it's 6
								d.knownSignals[6] = wire
								d.solved[i] = true
							}
						}
					}
				}
			}
			if len(d.knownSignals) == 10 {
				return
			}
		}
	}
}

func (d *display) setUniqeMappings(wire []string) {
	switch len(wire) {
	case 2:
		d.knownSignals[1] = wire
	case 4:
		d.knownSignals[4] = wire
	case 3:
		d.knownSignals[7] = wire
	case 7:
		d.knownSignals[8] = wire
	}
}

type display struct {
	wires        [][]string
	output       [][]string
	knownSignals map[int][]string
	solved       map[int]bool
}

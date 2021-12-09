package main

import (
	"log"
	"math"
)

func arrayContainsAll(a []string, b []string) bool {
	for _, s := range b {
		if !arrayContains(a, s) {
			return false
		}
	}
	return true
}

func arrayContainsExact(a []string, b []string) bool {
	return len(a) == len(b) && arrayContainsAll(a, b)
}

func stringToCharArray(s string) []string {
	a := make([]string, len(s))
	for i := range s {
		a[i] = string(s[i])
	}
	return a
}

func arrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (d *display) value() int {
	if len(d.knownSignals) != 10 {
		log.Println("Not solved yet")
		return 0
	}

	value := 0
	for i, o := range d.output {
		for v, s := range d.knownSignals {
			if arrayContainsExact(o, s) {
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
				if arrayContainsAll(wire, d.knownSignals[1]) {
					// It's 3
					d.knownSignals[3] = wire
					d.solved[i] = true
				} else if _, ok := d.knownSignals[6]; ok {
					if arrayContainsAll(d.knownSignals[6], wire) {
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
					if arrayContainsAll(wire, d.knownSignals[3]) {
						// It's 9
						d.knownSignals[9] = wire
						d.solved[i] = true
					} else {
						if _, ok := d.knownSignals[7]; ok {
							if arrayContainsAll(wire, d.knownSignals[7]) {
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

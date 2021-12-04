package main

func filterOxygen(bits [][]string, pos int) [][]string {
	common := mostCommon(bits, pos)

	var remainingbits [][]string
	for i := 0; i < len(bits); i++ {
		if bits[i][pos] == common {
			remainingbits = append(remainingbits, bits[i])
		}
	}
	if len(remainingbits) == 1 {
		return remainingbits
	}

	return filterOxygen(remainingbits, pos+1)
}

func filterCo2(bits [][]string, pos int) [][]string {

	common := leastCommon(bits, pos)

	var remainingbits [][]string
	for i := 0; i < len(bits); i++ {
		if bits[i][pos] == common {
			remainingbits = append(remainingbits, bits[i])
		}
	}
	if len(remainingbits) == 1 {
		return remainingbits
	}

	return filterCo2(remainingbits, pos+1)
}

func mostCommon(bits [][]string, pos int) string {
	zeroes, ones := count(bits, pos)
	if zeroes == ones {
		return "1"
	} else if zeroes > ones {
		return "0"
	} else {
		return "1"
	}
}

func leastCommon(bits [][]string, pos int) string {
	zeroes, ones := count(bits, pos)
	if zeroes > ones {
		return "1"
	} else {
		return "0"
	}
}

func count(bits [][]string, pos int) (int, int) {
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

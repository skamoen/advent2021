package main

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

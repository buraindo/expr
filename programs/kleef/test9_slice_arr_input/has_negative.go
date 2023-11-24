package main

func hasNegativeSlice(a []int) bool {
	for i := range a {
		if a[i] < 0 {
			return true
		}
	}
	return false
}

func hasNegativeArr(a [3]int) bool {
	for i := range a {
		if a[i] < 0 {
			return true
		}
	}
	return false
}

func hasNegativeVarargs(a ...int) bool {
	for i := range a {
		if a[i] < 0 {
			return true
		}
	}
	return false
}

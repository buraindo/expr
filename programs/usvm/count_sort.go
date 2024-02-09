package main

func count(data []int) []int {
	l := len(data)
	if l == 0 || l > 3 {
		return data
	}

	var aMin, aMax = data[0], data[0]
	for _, x := range data {
		if x < aMin {
			aMin = x
		}
		if x > aMax {
			aMax = x
		}
	}
	if aMax > 1000 || aMax < -1000 || aMin < -1000 || aMin > 1000 {
		return data
	}

	length := aMax - aMin + 1
	if length > 1000 {
		return data
	}
	if l < 3 {
		return data
	}

	allEqual := true
	for _, x := range data {
		if x != data[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return data
	}

	count := make([]int, length)
	for _, x := range data {
		count[x-aMin]++
	}

	z := 0
	out := make([]int, len(data))
	for i, c := range count {
		for c > 0 {
			out[z] = i + aMin
			z++
			c--
		}
	}

	return out
}

func main() {}

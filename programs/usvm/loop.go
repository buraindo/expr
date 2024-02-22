package main

func loop(i int) int {
	idx := i
	for idx < 10 {
		idx++
	}
	return 0
}

func loop2(i int) int {
	idx := i
	for idx < 10 {
		if idx < 5 {
			idx++
		} else {
			idx += 2
		}
	}
	return 0
}

func infinite(i int) int {
	for {
		j := i + 3
		k := j + i
		l := k + j + 4
		j = l + 2
	}
}

func inner(i, j int) int {
	sum := 0
	for k := 0; k < 100; k++ {
		for l := 0; l < 100; l++ {
			if k+l < i+j {
				sum += k + l
			}
		}
	}
	return sum
}

func collatz(i int) int {
	if i <= 0 || i >= 100 {
		return 0
	}

	j := i
	loopCount := 0

	for j != 1 {
		oldJ := j
		j = j*3 + 1
		if oldJ%2 == 0 {
			j = oldJ / 2
		}
		loopCount++
	}

	if loopCount == 17 {
		return 1
	}
	return 2
}

func sum(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func rangeMap(m map[int]int, target int) (int, bool) {
	if len(m) < 5 {
		return 1, false
	}
	for k, v := range m {
		if v == target {
			return k, true
		}
	}
	return -1, false
}

func rangeString(s string, target rune) int {
	state := 0
	if len(s) == 0 {
		state = 1
	}

	other := "hello"
	for i, v := range s {
		if s[i] == other[i] {
			state = 2
		}
		if v == target {
			state = 3
		}
	}
	return state
}

func main() {}

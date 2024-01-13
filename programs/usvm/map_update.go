package main

func update(m map[int]int, k, v int) int {
	vOld := m[k]
	m[k] = v
	if v > vOld {
		return v
	}
	return vOld
}

func main() {}

package main

func panicFunc(a int) int {
	if a > 0 {
		return 1
	}
	if a < 0 {
		return -1
	}
	panic(0)
}

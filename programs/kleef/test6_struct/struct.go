package main

type Input struct {
	A int
}

func getSign(i Input) int {
	if i.A > 0 {
		return 1
	}
	if i.A < 0 {
		return -1
	}
	return 0
}

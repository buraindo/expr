package main

func getSign(i *int) int {
	if *i == 0 {
		return 0
	}
	if *i > 0 {
		return 1
	}
	return -1
}


package main

func getSign(a int) int {
	if a == 0 {
		return 0
	}
	if a > 0 {
		return 1
	}
	return -1
}

func main() {
	print(getSign(250))
}

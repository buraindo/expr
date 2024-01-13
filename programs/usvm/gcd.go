package main

func gcd(a, b int64) int64 {
	if a == 0 || b == 0 {
		return -1
	}
	if a == 1 || b == 1 {
		return 1
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {}

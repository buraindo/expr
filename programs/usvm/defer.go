package main

func simple(a int) int {
	defer func() {
		a++
	}()

	return a
}

func main() {}

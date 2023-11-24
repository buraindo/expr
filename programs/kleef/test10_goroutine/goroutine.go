package main

func goroutine(a int) int {
	go func() {
		print("i do nothing")
	}()
	if a >= 0 {
		return 1
	}
	return 0
}

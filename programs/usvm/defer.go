package main

import "fmt"

func simple(a int) (b int) {
	defer func() {
		b++
		defer func() {
			b++
		}()
		defer func() {
			defer func() {
				defer func() {
					b++
				}()
				b++
			}()
			b++
		}()
	}()

	return a
}

func panicking() {
	panic("oh no")
}

func panicRecover(a int) (b int) {
	defer func() {
		if r := recover(); r != nil {
			b = 2
		}
	}()

	panicking()
	return 1000
}

func main() {
	fmt.Println(panicRecover(1))
}

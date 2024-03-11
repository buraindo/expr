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

func main() {}

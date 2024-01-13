package main

func loop(i int) int {
	for {
		j := i + 3
		k := j + i
		l := k + j + 4
		j = l + 2
	}
}

func main() {}

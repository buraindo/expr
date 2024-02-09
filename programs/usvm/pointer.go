package main

func pointer(i int) int {
	j := &i
	k := &j
	l := &k
	**k = 8
	**l = *k
	return ***l + 1
}

func main() {}

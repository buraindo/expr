package main

type iface1 interface {
	do() int
}

type iface2 interface {
	do() int
	other()
}

func change(i iface2) int {
	var j iface1 = i
	_ = j
	return j.do()
}

func main() {}

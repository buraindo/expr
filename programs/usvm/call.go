package main

import (
	"fmt"
	"sort"
)

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func stdSort(data []int) []int {
	if len(data) != 7 {
		return data
	}

	sort.Ints(data)
	return data
}

func swap(data sort.Interface, a, b int) (int, int) {
	if data.Less(b, a) {
		return b, a
	}
	return a, b
}

func callSwap(arr []int) (int, int) {
	return swap(sort.IntSlice(arr), 1, 3)
}

func callAnonClosure(i int) int {
	if i == 0 {
		return -1
	}
	f := func(j int) int {
		return i + j
	}
	return f(2)
}

type Person struct {
	Age int
}

func (p Person) getAge() int {
	return p.Age
}

func callMethod(i int) int {
	p := Person{Age: i + 2}
	return p.getAge()
}

func callBoundClosure(i int) int {
	p := Person{Age: i + 2}
	f := p.getAge
	return f()
}

type SuperInt int

func (s SuperInt) sqr() int {
	i := int(s)
	return i * i
}

func (s SuperInt) add(j int) int {
	i := int(s)
	return i + j
}

func callSuperIntSqr(i int) int {
	if i >= 0 && i <= 1 {
		return i
	}
	return SuperInt(i).sqr()
}

func callConstSuperIntAdd(i int) int {
	return SuperInt(2).add(i)
}

type Squarer interface {
	sqr() int
}

func callIntSquarerSqr(i int) int {
	if i >= 0 && i <= 1 {
		return i
	}
	return Squarer(SuperInt(i)).sqr()
}

type SquarerImpl struct {
	i int
}

func (s SquarerImpl) sqr() int {
	return s.i * s.i
}

func callStructSquarerSqr(i int) int {
	if i >= 0 && i <= 1 {
		return i
	}
	return Squarer(SquarerImpl{i: i}).sqr()
}

func main() {
	n := 2
	i := Squarer(SuperInt(n))
	j := Squarer(SuperInt(n))
	fmt.Println(i == j, i, j)
	n = 3
	fmt.Println(i == j, i, j)
}

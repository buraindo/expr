package main

import (
	"fmt"
	"strconv"
)

func returnString(i int) string {
	switch i {
	case 0:
		return "hello"
	case 1:
		return "world"
	default:
		return "hello, world!"
	}
}

func typeAssert(a int) int {
	var b interface{} = a
	c := b.(int64)
	return int(c)
}

// Алиас на примитив
type SuperInt int

// Реализуем интерфейс fmt.Stringer у этого алиаса
func (s SuperInt) String() string {
	return strconv.Itoa(int(s))
}

// Функция работает с интерфейсом
func printStringer(s fmt.Stringer) {
	fmt.Println(s.String())
}

// Функция кастит алиас к интерфейсу
func printSuperInt(s SuperInt) {
	printStringer(fmt.Stringer(s))
}

func printInt(i int) {
	printSuperInt(SuperInt(i))
}

func compareSuperInt(i, j SuperInt) {
	fmt.Println(i == j)
}

func changeSuperInt(i *SuperInt) {
	*i = 3
}

func main() {
	compareSuperInt(2, 2)
	printSuperInt(2)

	i := SuperInt(2)
	changeSuperInt(&i)
	fmt.Println(i)
}

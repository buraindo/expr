package main

type X interface {
	String() string
}

type testthing struct{}

type x chan int

func (t testthing) String() string {
	return "a test thing"
}

func main() {

}

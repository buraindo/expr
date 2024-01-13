package main

type Person struct {
	Age    int
	Gender bool
}

func isOld(p Person) bool {
	if p.Gender {
		return p.Age > 20
	}
	return p.Age > 40
}

func main() {}

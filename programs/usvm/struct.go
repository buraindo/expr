package main

type Person struct {
	Age    int
	Gender bool
}

func isOld(p Person) bool {
	if p.Gender {
		if p.Age > 20 {
			return true
		}
		return false
	}
	if p.Age > 40 {
		return true
	}
	return false
}

func setPerfectAge(p Person) int {
	if p.Gender {
		p.Age = 18
		return p.Age
	}
	p.Age = 25
	return p.Age
}

func main() {}

package armstrong

func pow(a, p int) int {
	r := a
	for i := 0; i < p-1; i++ {
		r *= a
	}
	return r
}

func IsArmstrong(n int) bool {
	var rightMost int
	var sum int = 0
	var tempNum int = n

	// to get the number of digits in the number
	length := 1
	tmp := n
	for tmp >= 10 {
		tmp /= 10
		length++
	}

	// get the right most digit and break the loop once all digits are iterated
	for tempNum > 0 {
		rightMost = tempNum % 10
		sum += pow(rightMost, length)

		// update the input digit minus the processed rightMost
		tempNum /= 10
	}

	return n == sum
}

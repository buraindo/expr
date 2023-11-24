package main

func SumOfProperDivisors(n uint) uint {
	var res = uint(0)
	if n > 1 {
		res = uint(1)
	}
	for curDivisor := uint(2); curDivisor*curDivisor <= n; curDivisor++ {
		if n%curDivisor == 0 {
			res += curDivisor
			if curDivisor*curDivisor != n {
				res += n / curDivisor
			}
		}
	}
	return res
}

func IsPerfectNumber(n uint) bool {
	return n > 0 && SumOfProperDivisors(n) == n
}

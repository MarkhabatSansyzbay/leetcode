package myFunctions

// 202. Happy Number

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	sum := 0
	m := make(map[int]bool)

	for n*n > 9 && !m[n] {
		m[n] = true
		sum = 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}
		n = sum
		if n == 1 {
			return true
		}
	}

	return false
}

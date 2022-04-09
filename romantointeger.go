package myFunctions

func romanToInt(s string) int {
	num := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if i < len(runes)-1 && roman(runes[i]) < roman(runes[i+1]) {
			num = num + roman(runes[i+1]) - roman(runes[i])
			i++
		} else {
			num += roman(runes[i])
		}
	}
	return num
}

func roman(r rune) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

package myFunctions

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		res += pow(26, len(columnTitle)-i-1) * int(columnTitle[i]-64)
	}
	return res
}

func pow(num, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res *= num
	}
	return res
}

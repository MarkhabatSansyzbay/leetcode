package myFunctions

// Given an integer x, return true if x is palindrome integer.

// An integer is a palindrome when it reads the same backward as forward.

// For example, 121 is a palindrome while 123 is not.

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var arr []rune
	for x > 0 {
		arr = append(arr, rune(x%10+48))
		x /= 10
	}
	var rev []rune
	for i := len(arr) - 1; i >= 0; i-- {
		rev = append(rev, arr[i])
	}
	if string(arr) == string(rev) {
		return true
	}
	return false
}

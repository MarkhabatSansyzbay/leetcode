package myFunctions

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	i := 0
	j := n - 1
	for p := n - 1; p >= 0; p-- {
		if abs(nums[i]) > abs(nums[j]) {
			result[p] = nums[i] * nums[i]
			i++
		} else {
			result[p] = nums[j] * nums[j]
			j--
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return num * (-1)
	}
	return num
}

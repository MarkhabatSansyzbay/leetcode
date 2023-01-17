package myFunctions

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

func rotate(nums []int, k int) {
	k = len(nums) - (k % len(nums))
	copy(nums, append(nums[k:], nums[:k]...))
}

package myFunctions

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// A subarray is a contiguous part of an array.

func maxSubArray(nums []int) int {
	sum := 0
	res := nums[0]
	max := nums[0]

	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		if sum < 0 {
			sum = 0
		}
		if nums[j] > max {
			max = nums[j]
		}
		if res < sum {
			res = sum
		}
	}

	if res == 0 {
		return max
	}
	return res
}

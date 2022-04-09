package myFunctions

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

func missingNumber(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	if nums[0] != 0 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
	}
	return nums[len(nums)-1] + 1
}

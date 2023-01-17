package myFunctions

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

func moveZeroes(nums []int) {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}

	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}

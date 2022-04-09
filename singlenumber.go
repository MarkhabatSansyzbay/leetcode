package myFunctions

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	flag := true
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] == nums[j] {
				flag = false
			}
		}
		if flag {
			return nums[i]
		}
		flag = true
	}
	return -1
}

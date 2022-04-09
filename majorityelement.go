package myFunctions

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

func majorityElement(nums []int) int {
	count := 1
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			res = nums[i]
		} else if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return res
}

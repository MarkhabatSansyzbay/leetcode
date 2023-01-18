package myFunctions

// 26. Remove Duplicates from Sorted Array

func removeDuplicates(nums []int) int {
	count := 1
	for i := 0; i < len(nums); i++ {
		if i != len(nums)-1 && nums[i] != nums[i+1] {
			nums[count] = nums[i+1]
			count++
		}
	}
	return count
}

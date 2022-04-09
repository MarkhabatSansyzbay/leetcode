package myFunctions

func removeDuplicates(nums []int) int {
	index := 1
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > temp {
			temp = nums[i]
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

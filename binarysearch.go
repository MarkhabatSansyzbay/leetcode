package myFunctions

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (end + start) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

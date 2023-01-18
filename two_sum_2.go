package myFunctions

// 167. Two Sum II - Input Array Is Sorted

/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.
*/

func twoSum2(numbers []int, target int) []int {
	res := make([]int, 2)
	n := len(numbers)
	m := make(map[int]int, n)

	for i := 0; i < n; i++ {
		j, ok := m[target-numbers[i]]
		if ok {
			res[0] = j
			res[1] = i + 1
			return res
		}
		m[numbers[i]] = i + 1
	}

	return res
}

package myFunctions

import "sort"

// 1200. Minimum Absolute Difference

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	var result [][]int
	for i := 0; i < len(arr)-1; i++ {
		res := make([]int, 2)
		if arr[i+1]-arr[i] < diff {
			diff = arr[i+1] - arr[i]
			result = nil
		}
		if arr[i+1]-arr[i] == diff {
			res[0] = arr[i]
			res[1] = arr[i+1]
			result = append(result, res)
		}
	}

	return result
}

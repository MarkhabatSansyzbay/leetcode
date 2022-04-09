package myFunctions

// You are climbing a staircase. It takes n steps to reach the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func climbStairs(n int) int {
	arr := []int{0, 1}
	for i := 0; i < n; i++ {
		arr = append(arr, arr[i]+arr[i+1])
	}
	return arr[len(arr)-1]
}

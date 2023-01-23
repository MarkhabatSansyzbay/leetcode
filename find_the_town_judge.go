package myFunctions

// 997. Find the Town Judge

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	m := map[int]int{}
	for _, w := range trust {
		m[w[1]]++
		m[w[0]]--
	}
	for key := range m {
		if m[key] == n-1 {
			return key
		}
	}

	return -1
}

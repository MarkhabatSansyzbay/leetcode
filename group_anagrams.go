package myFunctions

// 49. Group Anagrams

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		m[letterSum(strs[i])] = append(m[letterSum(strs[i])], strs[i])
	}

	var res [][]string
	for _, arr := range m {
		res = append(res, arr)
	}
	return res
}

func letterSum(s string) [26]int {
	var arr [26]int
	for i := range s {
		arr[int(s[i]-'a')] += 1
	}
	return arr
}

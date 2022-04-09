package myFunctions

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

func longestCommonPrefix(strs []string) string {
	prefix := ""
	var runes [][]rune
	for _, word := range strs {
		arr := []rune(word)
		runes = append(runes, arr)
	}

	shortest := strs[0]
	for i := 0; i < len(strs); i++ {
		if len(shortest) > len(strs[i]) {
			shortest = strs[i]
		}
	}

	flag := true
	for j := 0; j < len(shortest); j++ {
		for i := 0; i < len(runes); i++ {
			if i != len(runes)-1 && runes[i][j] != runes[i+1][j] {
				flag = false
			}
		}
		if flag {
			prefix += string(shortest[j])
		} else {
			break
		}
		flag = true
	}

	return prefix
}

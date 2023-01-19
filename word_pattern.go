package myFunctions

// 290. Word Pattern

func wordPattern(pattern string, s string) bool {
	words := split(s)
	if len(pattern) != len(words) {
		return false
	}

	match := make(map[rune]string)
	match2 := make(map[string]rune)

	for i, v := range pattern {
		word, ok := match[v]
		pattern, ok2 := match2[words[i]]
		if ok && word != words[i] || ok2 && pattern != v {
			return false
		}
		match[v] = words[i]
		match2[words[i]] = v
	}

	return true
}

func split(s string) []string {
	var res []string
	word := ""

	for _, v := range s {
		if v != ' ' {
			word += string(v)
		} else if v == ' ' {
			res = append(res, word)
			word = ""
		}
	}

	if word != "" {
		res = append(res, word)
	}

	return res
}

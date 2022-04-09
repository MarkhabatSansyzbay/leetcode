package myFunctions

// We define the usage of capitals in a word to be right when one of the following cases holds:

// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Given a string word, return true if the usage of capitals in it is right.

func detectCapitalUse(word string) bool {
	upper := false
	lower := false
	first := true
	if len(word) <= 1 {
		return true
	} else if word[0] < 65 || word[0] > 90 {
		first = false
	}
	for _, r := range word[1:] {
		if r >= 'A' && r <= 'Z' {
			upper = true
		} else {
			lower = true
		}
	}
	return !(upper && lower) && first || !first && lower && !upper
}

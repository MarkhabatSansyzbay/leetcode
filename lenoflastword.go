package myFunctions

// Given a string s consisting of some words separated by some number of spaces, return the length of the last word in the string.

func lengthOfLastWord(s string) int {
	word := ""
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			word += string(s[i])
		} else if word != "" && string(s[i]) == " " {
			break
		}
	}
	return len(word)
}

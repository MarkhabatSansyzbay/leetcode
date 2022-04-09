package myFunctions

func isPalindromeString(s string) bool {
	var runes []rune
	for i := range s {
		if s[i] >= 97 && s[i] <= 122 || s[i] >= 48 && s[i] <= 57 {
			runes = append(runes, rune(s[i]))
		} else if s[i] >= 65 && s[i] <= 90 {
			runes = append(runes, rune(s[i]+32))
		}
	}
	for i := 0; i < len(runes); i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}

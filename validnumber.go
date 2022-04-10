package myFunctions

func isNumber(s string) bool {
	s1 := ""
	s2 := ""
	temp := ""
	flagE := false
	for i := range s {
		temp += string(s[i])
		if (s[i] == 'e' || s[i] == 'E') && !flagE {
			flagE = true
			s1 = temp[:i]
			temp = ""
		} else if i == len(s)-1 {
			s2 = temp
		}
	}

	if flagE {
		if isDecimal(s1) && isInteger(s2) {
			return true
		}
	} else {
		if isInteger(s) || isDecimal(s) {
			return true
		}
	}

	return false
}

func isDecimal(s string) bool {
	if s == "" {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if s == "" {
			return false
		}
	}

	flag := true
	for i := 0; i < len(s); i++ {
		if s[i] == '.' && flag && len(s) > 1 {
			flag = false
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}

func isInteger(s string) bool {
	if s == "" {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if s == "" {
			return false
		}
	}

	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}

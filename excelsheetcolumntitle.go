package myFunctions

// Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

func convertToTitle(columnNumber int) string {
	var arr []rune
	for columnNumber > 0 {
		arr = append(arr, rune((columnNumber-1)%26+65))
		columnNumber = (columnNumber - 1) / 26
	}
	res := make([]rune, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[len(arr)-i-1]
	}
	return string(res)
}

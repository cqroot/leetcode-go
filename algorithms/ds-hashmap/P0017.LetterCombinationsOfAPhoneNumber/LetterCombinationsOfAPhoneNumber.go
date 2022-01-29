package p0017lettercombinationsofaphonenumber

import "fmt"

func letterCombinations(digits string) []string {
	digitsMap := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	if len(digits) == 0 {
		return []string{}
	}

	letterStr := digitsMap[int(digits[0]-'0')]

	lettersSlice := []string{}
	lettersSlice = letterCombinations(digits[1:])
	if len(lettersSlice) == 0 {
		lettersSlice = []string{""}
	}

	ans := []string{}
	for _, letter := range letterStr {
		for _, letters := range lettersSlice {
			ans = append(ans, fmt.Sprintf("%c%s", letter, letters))
		}
	}

	return ans
}

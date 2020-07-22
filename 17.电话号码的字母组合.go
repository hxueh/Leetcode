/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

func helper(digits string, res string, resArray *[]string, phoneCombMap map[byte]string) {
	if len(digits) == 0 {
		*resArray = append(*resArray, res)
	} else {
		valueString, digits := phoneCombMap[digits[0]], digits[1:]
		for _, value := range valueString {
			helper(digits, res + string(value), resArray, phoneCombMap)
		}
	}
}

func letterCombinations(digits string) []string {
	phoneCombMap := map[byte]string{
		'2': "abc",
    	'3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
	}
	resArray := make([]string, 0)
	if len(digits) > 0 {
		helper(digits, "", &resArray, phoneCombMap)
	}
	return resArray
}

// Time complexity O(M*N)
// Space complexity O(N)

// @lc code=end

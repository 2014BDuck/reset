// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lower-case English letters.

func longestCommonPrefix(strs []string) string {
	idx := 0
	result := ""

	for {
		if idx >= len(strs[0]) {
			return result
		}
		cur := strs[0][idx]
		for _, s := range strs {

			if idx >= len(s) {
				return result
			}

			if s[idx] != cur {
				return result
			}
		}

		result += string(cur)
		idx++
	}
	return result
}
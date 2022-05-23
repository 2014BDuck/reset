// Given a string s, return the longest palindromic substring in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.

func longestPalindrome(s string) string {
	result := ""
	n := len(s)

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > len(result) {
				result = s[l : r+1]
			}
			l--
			r++
		}
	}

	for i := 0; i < n-1; i++ {
		l, r := i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > len(result) {
				result = s[l : r+1]
			}
			l--
			r++
		}
	}
	return result
}
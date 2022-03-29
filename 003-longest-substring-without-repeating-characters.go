// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:

// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.

func lengthOfLongestSubstring(s string) int {
    start, maxLength := 0, 0
    usedChar := map[byte]int{}

    for i := 0; i < len(s); i++ {
        idx, ok := usedChar[s[i]]
        if ok && start <= idx {
            start = idx + 1
        } else {
            if maxLength < i-start+1 {
                maxLength = i - start + 1
            }
        }

        usedChar[s[i]] = i
    }

    return maxLength
}
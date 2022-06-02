// Given a string s, return the number of palindromic substrings in it.

// A string is a palindrome when it reads the same backward as forward.

// A substring is a contiguous sequence of characters within the string.

// Example 1:

// Input: s = "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
// Example 2:

// Input: s = "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

// Constraints:

// 1 <= s.length <= 1000
// s consists of lowercase English letters.

func countSubstrings(s string) int {
    n := len(s)

    result := 0
    for i := 0; i < n; i++ {
        // odd
        l, r := i, i
        for l >= 0 && r < n {
            if s[l] == s[r] {
                result++
                l--
                r++
                continue
            }
            break
        }

        // even
        l, r = i, i+1
        for l >= 0 && r < n {
            if s[l] == s[r] {
                result++
                l--
                r++
                continue
            }
            break
        }
    }
    return result
}
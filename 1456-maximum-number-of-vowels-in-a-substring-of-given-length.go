// Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

// Example 1:

// Input: s = "abciiidef", k = 3
// Output: 3
// Explanation: The substring "iii" contains 3 vowel letters.
// Example 2:

// Input: s = "aeiou", k = 2
// Output: 2
// Explanation: Any substring of length 2 contains 2 vowels.
// Example 3:

// Input: s = "leetcode", k = 3
// Output: 2
// Explanation: "lee", "eet" and "ode" contain 2 vowels.

// Constraints:

// 1 <= s.length <= 105
// s consists of lowercase English letters.
// 1 <= k <= s.length

var m = map[byte]bool{
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func maxVowels(s string, k int) int {
    n := len(s)
    counter := 0

    l, r := 0, 0

    for r-l < k {
        if m[s[r]] {
            counter++
        }
        r++
    }
    result := counter
    r--

    for r < n-1 {
        // r and l move 1 idx forward
        if m[s[l]] {
            counter--
        }
        l++

        r++
        if m[s[r]] {
            counter++
        }
        if counter == k {
            return k
        }
        result = max(result, counter)
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

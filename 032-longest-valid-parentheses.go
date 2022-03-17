// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

// Example 1:

// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".
// Example 2:

// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".
// Example 3:

// Input: s = ""
// Output: 0

// Constraints:

// 0 <= s.length <= 3 * 104
// s[i] is '(', or ')'.

func longestValidParentheses(s string) int {
    result := 0
    if len(s) < 2 {
        return result
    }

    stack := []int{-1}

    for i := range s {
        if len(stack) > 1 && s[i] == ')' && s[stack[len(stack)-1]] == '(' {
            stack = stack[:len(stack)-1]
            result = max(result, i-stack[len(stack)-1])
        } else {
            stack = append(stack, i)
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

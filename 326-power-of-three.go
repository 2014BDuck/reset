// Given an integer n, return true if it is a power of three. Otherwise, return false.

// An integer n is a power of three, if there exists an integer x such that n == 3x.

// Example 1:

// Input: n = 27
// Output: true
// Example 2:

// Input: n = 0
// Output: false
// Example 3:

// Input: n = 9
// Output: true

// Constraints:

// -231 <= n <= 231 - 1

// Follow up: Could you solve it without loops/recursion?

func isPowerOfThree(n int) bool {
    if n == 1 {
        return true
    }
    m := n / 3
    if m*3 != n {
        return false
    }
    if m < 1 {
        return false
    } else if m == 1 {
        return true
    } else {
        return isPowerOfThree(m)
    }
}

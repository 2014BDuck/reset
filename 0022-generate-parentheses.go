// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// Example 1:

// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// Example 2:

// Input: n = 1
// Output: ["()"]

// Constraints:

// 1 <= n <= 8

func generateParenthesis(n int) []string {
    m := map[string]bool{"": true}

    for i := 0; i < n; i++ {
        t := map[string]bool{}

        for k, _ := range m {
            for j := 0; j < 2*i+1; j++ {
                t[k[:j]+"()"+k[j:]] = true
            }
        }
        m = t
    }

    result := []string{}
    for k, _ := range m {
        result = append(result, k)
    }
    return result
}
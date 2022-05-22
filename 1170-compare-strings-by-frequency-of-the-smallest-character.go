// Let the function f(s) be the frequency of the lexicographically smallest character in a non-empty string s. For example, if s = "dcce" then f(s) = 2 because the lexicographically smallest character is 'c', which has a frequency of 2.

// You are given an array of strings words and another array of query strings queries. For each query queries[i], count the number of words in words such that f(queries[i]) < f(W) for each W in words.

// Return an integer array answer, where each answer[i] is the answer to the ith query.

// Example 1:

// Input: queries = ["cbd"], words = ["zaaaz"]
// Output: [1]
// Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").
// Example 2:

// Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
// Output: [1,2]
// Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").

// Constraints:

// 1 <= queries.length <= 2000
// 1 <= words.length <= 2000
// 1 <= queries[i].length, words[i].length <= 10
// queries[i][j], words[i][j] consist of lowercase English letters.

func numSmallerByFrequency(queries []string, words []string) []int {
    n := len(words)
    memory := make([]int, n, n)
    for i := range words {
        memory[i] = calFreq(words[i])
    }

    sort.Ints(memory)
    result := make([]int, len(queries), len(queries))
    for i := range queries {
        freq := calFreq(queries[i])
        for j := 0; j < len(memory); j++ {
            if memory[j] > freq {
                result[i] = n - j
                break
            }
        }
    }
    return result
}

func calFreq(a string) int {
    small := rune('z')
    count := 0

    for _, c := range a {
        if c == small {
            count++
        } else if c < small {
            small = c
            count = 1
        }
    }
    return count
}
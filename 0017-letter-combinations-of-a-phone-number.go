// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example 1:

// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Example 2:

// Input: digits = ""
// Output: []
// Example 3:

// Input: digits = "2"
// Output: ["a","b","c"]

// Constraints:

// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].

func letterCombinations(digits string) []string {
    m := map[byte][]string{
        '0': []string{"0"},
        '1': []string{"1"},
        '2': []string{"a", "b", "c"},
        '3': []string{"d", "e", "f"},
        '4': []string{"g", "h", "i"},
        '5': []string{"j", "k", "l"},
        '6': []string{"m", "n", "o"},
        '7': []string{"p", "q", "r", "s"},
        '8': []string{"t", "u", "v"},
        '9': []string{"w", "x", "y", "z"},
    }

    result := []string{""}
    for _, c := range digits {
        tmp := []string{}
        for _, s := range result {
            for _, p := range m[byte(c)] {
                tmp = append(tmp, s+p)
            }
        }
        result = tmp
    }
    if len(result) == 1 && result[0] == "" {
        result = []string{}
    }
    return result
}
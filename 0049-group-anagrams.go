// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

// Constraints:

// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.

func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}
    l := make([]string, len(strs), len(strs))
    for i, v := range strs {
        l[i] = SortString(v)
    }

    for i, v := range l {
        m[v] = append(m[v], strs[i])
    }
    result := [][]string{}
    for _, v := range m {
        result = append(result, v)
    }
    return result
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
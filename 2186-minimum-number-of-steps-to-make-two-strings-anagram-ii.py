# You are given two strings s and t. In one step, you can append any character to either s or t.

# Return the minimum number of steps to make s and t anagrams of each other.

# An anagram of a string is a string that contains the same characters with a different (or the same) ordering.

# Example 1:

# Input: s = "leetcode", t = "coats"
# Output: 7
# Explanation: 
# - In 2 steps, we can append the letters in "as" onto s = "leetcode", forming s = "leetcodeas".
# - In 5 steps, we can append the letters in "leede" onto t = "coats", forming t = "coatsleede".
# "leetcodeas" and "coatsleede" are now anagrams of each other.
# We used a total of 2 + 5 = 7 steps.
# It can be shown that there is no way to make them anagrams of each other with less than 7 steps.
# Example 2:

# Input: s = "night", t = "thing"
# Output: 0
# Explanation: The given strings are already anagrams of each other. Thus, we do not need any further steps.

# Constraints:

# 1 <= s.length, t.length <= 2 * 105
# s and t consist of lowercase English letters.

from collections import defaultdict
class Solution(object):
    def minSteps(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        s1 = defaultdict(int)
        s2 = defaultdict(int)
        for each in s:
            s1[each] = s1[each] +1
        for each in t:
            s2[each] = s2[each] +1
        
        s3, s4 = s1.copy(), s2.copy()
        for k, v in s1.items():
            if k in s2:
                s4[k] -= min(s2[k], s1[k])
        
        for k, v in s2.items():
            if k in s1:
                s3[k] -= min(s1[k], s2[k])
        
        c = 0
        for k, v in s3.items():
            c += v
        for k, v in s4.items():
            c += v

        return c

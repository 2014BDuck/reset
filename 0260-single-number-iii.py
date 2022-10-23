from collections import Counter

class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = Counter(nums)
        result = []
        for k,v in n.items():
            if v == 1:
                result.append(k)
                
        return result

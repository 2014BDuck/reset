class Solution(object):
    def countDistinctIntegers(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        reverse = []
        for num in nums:
            reverse.append(int(str(num)[::-1]))
            
        return len(set(nums+reverse))

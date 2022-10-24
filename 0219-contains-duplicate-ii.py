
class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        m = {}
        for i in range(len(nums)):
            if nums[i] in m and i - m[nums[i]] <= k:
                return True
            m[nums[i]] = i
        return False

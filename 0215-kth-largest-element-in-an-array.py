# Given an integer array nums and an integer k, return the kth largest element in the array.

# Note that it is the kth largest element in the sorted order, not the kth distinct element.

# You must solve it in O(n) time complexity.

# Example 1:

# Input: nums = [3,2,1,5,6,4], k = 2
# Output: 5
# Example 2:

# Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
# Output: 4
 

# Constraints:

# 1 <= k <= nums.length <= 105
# -104 <= nums[i] <= 104

class Solution(object):
    def findKthLargest(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        if not nums:
            return 
    
        pivot = nums[len(nums) // 2]
        left = [x for x in nums if x > pivot]
        mid = l = [x for x in nums if x == pivot]
        right = [x for x in nums if x < pivot]
        
        l, m = len(left), len(mid)
        
        if k <= l:
            return self.findKthLargest(left, k)
        elif k > l + m:
            return self.findKthLargest(right, k-l-m)
        else:
            return mid[0]

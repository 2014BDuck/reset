# You are given an array of strings names, and an array heights that consists of distinct positive integers. Both arrays are of length n.

# For each index i, names[i] and heights[i] denote the name and height of the ith person.

# Return names sorted in descending order by the people's heights.

# Example 1:

# Input: names = ["Mary","John","Emma"], heights = [180,165,170]
# Output: ["Mary","Emma","John"]
# Explanation: Mary is the tallest, followed by Emma and John.
# Example 2:

# Input: names = ["Alice","Bob","Bob"], heights = [155,185,150]
# Output: ["Bob","Alice","Bob"]
# Explanation: The first Bob is the tallest, followed by Alice and the second Bob.

# Constraints:

# n == names.length == heights.length
# 1 <= n <= 103
# 1 <= names[i].length <= 20
# 1 <= heights[i] <= 105
# names[i] consists of lower and upper case English letters.
# All the values of heights are distinct.

class Solution(object):
    def sortPeople(self, names, heights):
        """
        :type names: List[str]
        :type heights: List[int]
        :rtype: List[str]
        """
        quick_sort(heights, names)
        return names[::-1]
        
        
def quick_sort(A, B):
    quick_sort2(A, B, 0, len(A)-1)

def quick_sort2(A, B, low, hi):
    # 递归方法
    if hi > low:
        p = partition(A, B, low, hi)
        quick_sort2(A, B, low, p-1)
        quick_sort2(A, B, p+1, hi)

def partition(A, B, low, hi):
    # 获取轴
    # [轴, 小于轴, 小于轴, 小于轴, 大于轴, 大于轴, 大于轴, 大于轴]
    # [小于轴, 小于轴, 小于轴, 轴, 大于轴, 大于轴, 大于轴, 大于轴]
    # 返回轴
    privot = get_privot(A, low, hi)
    privot_value = A[privot]
    A[low], A[privot] = privot_value, A[low]
    B[low], B[privot] = B[privot], B[low]
    border = low
    for i in range(low+1, hi+1):
        if A[i] < privot_value:
            border += 1
            A[i], A[border] = A[border], A[i]
            B[i], B[border] = B[border], B[i]

    A[low], A[border] = A[border], A[low]
    B[low], B[border] = B[border], B[low]
    return border

def get_privot(A, low, hi):
    # 选取一个位置为轴
    # 若不选取中间值为轴 最差O(n^2)
    mid = (hi + low) // 2
    s = sorted([A[hi], A[low], A[mid]])
    if s[1] == A[hi]:
        return hi
    elif s[1] == A[low]:
        return low
    else:
        return mid
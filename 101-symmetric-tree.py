# Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

 

# Example 1:


# Input: root = [1,2,2,3,4,4,3]
# Output: true
# Example 2:


# Input: root = [1,2,2,null,3,null,3]
# Output: false
 

# Constraints:

# The number of nodes in the tree is in the range [1, 1000].
# -100 <= Node.val <= 100
 

# Follow up: Could you solve it both recursively and iteratively?

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        
        def dfs(l, r):
            if l and r:
                if l.val != r.val:
                    return False
                return dfs(l.left, r.right) and dfs(l.right, r.left)
            elif not l and not r:
                return True
            else:
                return False
        
        if not root:
            return True
        
        return dfs(root.left, root.right)
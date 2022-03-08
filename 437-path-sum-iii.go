// Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

// The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

// Example 1:

// Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// Output: 3
// Explanation: The paths that sum to 8 are shown.
// Example 2:

// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: 3

// Constraints:

// The number of nodes in the tree is in the range [0, 1000].
// -109 <= Node.val <= 109
// -1000 <= targetSum <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return helper(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func helper(root *TreeNode, sum int) int {
	count := 0
	if root == nil {
		return 0
	}
	if root.Val == sum {
		count++
	}
	return count + helper(root.Left, sum-root.Val) + helper(root.Right, sum-root.Val)
}
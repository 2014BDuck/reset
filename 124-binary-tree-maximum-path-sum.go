// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

// The path sum of a path is the sum of the node's values in the path.

// Given the root of a binary tree, return the maximum path sum of any non-empty path.

// Example 1:

// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
// Example 2:

// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

// Constraints:

// The number of nodes in the tree is in the range [1, 3 * 104].
// -1000 <= Node.val <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	l, r, m := dfs(root)

	return max(max(l, r), m)
}

func dfs(node *TreeNode) (int, int, int) {
	if node == nil {
		return -999999999999, -999999999999, -999999999999
	}

	l1, r1, m1 := dfs(node.Left)
	l2, r2, m2 := dfs(node.Right)

	l := max(max(l1, r1)+node.Val, node.Val)
	r := max(max(l2, r2)+node.Val, node.Val)
	m := max(max(max(m1, m2), max(l1, r1)+node.Val+max(l2, r2)), node.Val)
	m = max(max(l, r), m)
	return l, r, m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
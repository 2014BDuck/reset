// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

// Example 1:

// Input: root = [3,1,4,null,2], k = 1
// Output: 1
// Example 2:

// Input: root = [5,3,6,2,4,null,null,1], k = 3
// Output: 3

// Constraints:

// The number of nodes in the tree is n.
// 1 <= k <= n <= 104
// 0 <= Node.val <= 104

// Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	result, _ := dfs(root, 0, k)
	return result.Val
}

func dfs(node *TreeNode, count, k int) (*TreeNode, int) {
	if node == nil {
		return nil, count
	}
	result, count := dfs(node.Left, count, k)
	if result != nil {
		return result, 0
	}
	count++
	if count == k {
		return node, 0
	}
	result, count = dfs(node.Right, count, k)
	if result != nil {
		return result, 0
	}
	return nil, count
}
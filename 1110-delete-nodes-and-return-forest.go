// Given the root of a binary tree, each node in the tree has a distinct value.

// After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).

// Return the roots of the trees in the remaining forest. You may return the result in any order.

// Example 1:

// Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
// Output: [[1,2,null,4],[6],[7]]
// Example 2:

// Input: root = [1,2,4,null,3], to_delete = [3]
// Output: [[1,2,4]]

// Constraints:

// The number of nodes in the given tree is at most 1000.
// Each node has a distinct value between 1 and 1000.
// to_delete.length <= 1000
// to_delete contains distinct values between 1 and 1000.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	m := map[int]bool{}
	for _, v := range to_delete {
		m[v] = true
	}

	result := make([]*TreeNode, 0, 1000)
	if root == nil {
		return result
	}

	result = append(result, root)
	result = del(root, m, result)
	if m[root.Val] {
		result = result[1:]
	}
	return result
}

func del(node *TreeNode, m map[int]bool, result []*TreeNode) []*TreeNode {
	if node == nil {
		return result
	}

	if m[node.Val] {
		// need to remove
		if node.Left != nil && !m[node.Left.Val] {
			result = append(result, node.Left)
		}
		if node.Right != nil && !m[node.Right.Val] {
			result = append(result, node.Right)
		}
	}

	result = del(node.Left, m, result)
	result = del(node.Right, m, result)
	if node.Left != nil && m[node.Left.Val] {
		node.Left = nil
	}
	if node.Right != nil && m[node.Right.Val] {
		node.Right = nil
	}
	return result
}
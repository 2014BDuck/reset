// Given the root of a binary tree, return the length of the longest path, where each node in the path has the same value. This path may or may not pass through the root.

// The length of the path between two nodes is represented by the number of edges between them.

// Example 1:

// Input: root = [5,4,5,1,1,5]
// Output: 2
// Example 2:

// Input: root = [1,4,5,4,4,5]
// Output: 2

// Constraints:

// The number of nodes in the tree is in the range [0, 104].
// -1000 <= Node.val <= 1000
// The depth of the tree will not exceed 1000.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result = 1

func longestUnivaluePath(root *TreeNode) int {
    result = 1
    dfs(root)
    return result - 1
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := dfs(root.Left)
    r := dfs(root.Right)

    if root.Left != nil && root.Right != nil && root.Left.Val == root.Val && root.Right.Val == root.Val {
        if l+r+1 > result {
            result = l + r + 1
        }
    }

    l0 := 0
    r0 := 0
    if root.Left != nil && root.Left.Val == root.Val {
        if l+1 > result {
            result = l + 1
        }
        l0 = l + 1
    }
    if root.Right != nil && root.Right.Val == root.Val {
        if r+1 > result {
            result = r + 1
        }
        r0 = r + 1
    }

    if l0 != 0 || r0 != 0 {
        if l0 > r0 {
            return l0
        }
        return r0
    }
    return 1
}
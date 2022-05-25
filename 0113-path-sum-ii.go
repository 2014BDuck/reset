// Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

// A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

// Example 1:

// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: [[5,4,11,2],[5,8,4,5]]
// Explanation: There are two paths whose sum equals targetSum:
// 5 + 4 + 11 + 2 = 22
// 5 + 8 + 4 + 5 = 22
// Example 2:

// Input: root = [1,2,3], targetSum = 5
// Output: []
// Example 3:

// Input: root = [1,2], targetSum = 0
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 5000].
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return [][]int{}
    }
    result = make([][]int, 0, 5000)

    dfs(root, targetSum, 0, []int{})

    return result
}

func dfs(node *TreeNode, targetSum, cur int, curPath []int) {
    if node == nil {
        return
    }
    cur += node.Val
    curPath = append(curPath, node.Val)
    if node.Left == nil && node.Right == nil {
        // leaf
        if cur == targetSum {
            dst := make([]int, len(curPath))
            copy(dst, curPath)
            result = append(result, dst)
            return
        }
    }

    dfs(node.Left, targetSum, cur, curPath)
    dfs(node.Right, targetSum, cur, curPath)
    return
}
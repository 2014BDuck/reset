// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]
// Example 2:

// Input: root = [1]
// Output: [[1]]
// Example 3:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 2000].
// -1000 <= Node.val <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	level := []*TreeNode{root}

	for len(level) > 0 {
		newLevel := []*TreeNode{}
		levelVal := []int{}
		for i := range level {
			levelVal = append(levelVal, level[i].Val)
			if level[i].Left != nil {
				newLevel = append(newLevel, level[i].Left)
			}
			if level[i].Right != nil {
				newLevel = append(newLevel, level[i].Right)
			}
		}
		level = newLevel
		result = append(result, levelVal)
	}
	return result

}
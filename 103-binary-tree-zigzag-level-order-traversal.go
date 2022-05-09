// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]
// Example 2:

// Input: root = [1]
// Output: [[1]]
// Example 3:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 2000].
// -100 <= Node.val <= 100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	level := []*TreeNode{root}

	result := [][]int{}
	flag := true
	for len(level) > 0 {
		var tmp []*TreeNode
		var line []int

		for i := range level {
			line = append(line, level[i].Val)

			if level[i].Left != nil {
				tmp = append(tmp, level[i].Left)
			}

			if level[i].Right != nil {
				tmp = append(tmp, level[i].Right)
			}
		}
		if !flag {
			reverseSlice(line)
		}
		result = append(result, line)

		flag = !flag

		level = tmp
	}
	return result
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
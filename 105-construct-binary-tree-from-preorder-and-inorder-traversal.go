// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

// Example 1:

// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]
// Example 2:

// Input: preorder = [-1], inorder = [-1]
// Output: [-1]

// Constraints:

// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder and inorder consist of unique values.
// Each value of inorder also appears in preorder.
// preorder is guaranteed to be the preorder traversal of the tree.
// inorder is guaranteed to be the inorder traversal of the tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var globalIdx = 0
var inorderMap map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	globalIdx = 0
	inorderMap = map[int]int{}
	for i, v := range inorder {
		inorderMap[v] = i
	}

	return helper(preorder, 0, len(inorder)-1)
}

func helper(preorder []int, left, right int) *TreeNode {
	if right < left {
		return nil
	}

	rootVal := preorder[globalIdx]
	root := &TreeNode{
		Val: rootVal,
	}
	globalIdx++
	root.Left = helper(preorder, left, inorderMap[rootVal]-1)
	root.Right = helper(preorder, inorderMap[rootVal]+1, right)
	return root
}
// Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

// Example 1:

// Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// Output: [3,9,20,null,null,15,7]
// Example 2:

// Input: inorder = [-1], postorder = [-1]
// Output: [-1]

// Constraints:

// 1 <= inorder.length <= 3000
// postorder.length == inorder.length
// -3000 <= inorder[i], postorder[i] <= 3000
// inorder and postorder consist of unique values.
// Each value of postorder also appears in inorder.
// inorder is guaranteed to be the inorder traversal of the tree.
// postorder is guaranteed to be the postorder traversal of the tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    node, _ := dfs(inorder, postorder)
    return node
}

func dfs(inorder []int, postorder []int) (*TreeNode, []int) {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil, postorder
    }

    root := &TreeNode{Val: postorder[len(postorder)-1]}
    postorder = postorder[:len(postorder)-1]
    idx := -1
    for i := range inorder {
        if inorder[i] == root.Val {
            idx = i
            break
        }
    }

    root.Right, postorder = dfs(inorder[idx+1:], postorder)
    root.Left, postorder = dfs(inorder[:idx], postorder)
    return root, postorder
}
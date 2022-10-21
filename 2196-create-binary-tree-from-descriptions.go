/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    m := make(map[int]*TreeNode)
    r := make(map[*TreeNode]bool)
    
    for i := range descriptions {
        parentNode := m[descriptions[i][0]]
        if parentNode == nil {
            parentNode = &TreeNode{Val: descriptions[i][0]}
        }
        
        childNode := m[descriptions[i][1]]
        if childNode == nil {
            childNode = &TreeNode{Val: descriptions[i][1]}
        }
        
        if descriptions[i][2] == 1 {
            parentNode.Left = childNode
        }else {
            parentNode.Right = childNode
        }
        
        m[parentNode.Val] = parentNode
        m[childNode.Val] = childNode
        
        r[parentNode] = true
        r[childNode] = true
    }
    
    for _, c := range m {
        if c.Left != nil {
            delete(r, c.Left)
        }
        if c.Right != nil {
            delete(r, c.Right)
        }
    }
    
    for root, _ := range r {
        return root
    }
    return nil
}

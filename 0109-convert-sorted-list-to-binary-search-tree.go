// Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

// Example 1:

// Input: head = [-10,-3,0,5,9]
// Output: [0,-3,9,-10,null,5]
// Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
// Example 2:

// Input: head = []
// Output: []

// Constraints:

// The number of nodes in head is in the range [0, 2 * 104].
// -105 <= Node.val <= 105

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }

    if head.Next == nil {
        return &TreeNode{Val: head.Val}
    }

    slow, fast := head, head.Next.Next

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    mid := slow.Next

    // unlink
    slow.Next = nil

    root := &TreeNode{Val: mid.Val}
    root.Left = sortedListToBST(head)
    root.Right = sortedListToBST(mid.Next)
    return root
}
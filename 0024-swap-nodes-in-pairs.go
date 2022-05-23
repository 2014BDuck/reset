// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

// Example 1:

// Input: head = [1,2,3,4]
// Output: [2,1,4,3]
// Example 2:

// Input: head = []
// Output: []
// Example 3:

// Input: head = [1]
// Output: [1]

// Constraints:

// The number of nodes in the list is in the range [0, 100].
// 0 <= Node.val <= 100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	node := head
	headChanged := false
	var prev *ListNode

	for node != nil && node.Next != nil {
		tmp := node.Next.Next
		if !headChanged {
			head = node.Next
			headChanged = true
		}
		if prev != nil {
			prev.Next = node.Next
		}
		node.Next.Next = node
		node.Next = tmp
		prev = node
		node = tmp
	}

	return head
}
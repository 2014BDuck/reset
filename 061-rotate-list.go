// Given the head of a linked list, rotate the list to the right by k places.

// Example 1:

// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]
// Example 2:

// Input: head = [0,1,2], k = 4
// Output: [2,0,1]

// Constraints:

// The number of nodes in the list is in the range [0, 500].
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 109

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	n := 0

	if head == nil {
		return nil
	}

	node := head
	for node != nil {
		n++
		node = node.Next
	}

	k = k % n
	if k == 0 {
		return head
	}

	node = head
	var last *ListNode
	for i := 0; i <= n-k-1; i++ {
		last = node
		node = node.Next
	}

	if last != nil {
		last.Next = nil
	}

	newHead := node

	node = newHead
	for node.Next != nil {
		node = node.Next
	}

	node.Next = head

	return newHead
}
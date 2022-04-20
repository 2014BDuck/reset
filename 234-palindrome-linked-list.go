// Given the head of a singly linked list, return true if it is a palindrome.

// Example 1:

// Input: head = [1,2,2,1]
// Output: true
// Example 2:

// Input: head = [1,2]
// Output: false

// Constraints:

// The number of nodes in the list is in the range [1, 105].
// 0 <= Node.val <= 9

// Follow up: Could you do it in O(n) time and O(1) space?

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    // copy list
    dummy := &ListNode{}
    node := head
    cpNode := dummy

    for node != nil {
        newNode := &ListNode{Val: node.Val}
        cpNode.Next = newNode

        node = node.Next
        cpNode = cpNode.Next
    }

    // reverse list
    var prev *ListNode
    node = dummy.Next

    for node != nil {
        tmp := node.Next
        node.Next = prev
        prev = node
        node = tmp
    }

    // prev -> nil vs head -> nil
    fmt.Println(head)
    fmt.Println(prev)
    for prev != nil {
        if prev.Val != head.Val {
            return false
        }
        prev = prev.Next
        head = head.Next
    }

    if prev == nil && head == nil {
        return true
    }
    return false
}
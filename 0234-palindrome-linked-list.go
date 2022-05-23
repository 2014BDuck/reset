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
    n := 0

    node := head
    for node != nil {
        n++
        node = node.Next
    }

    mid := n / 2

    node = head
    var prev *ListNode
    for i := 0; i < mid; i++ {
        tmp := node.Next
        node.Next = prev
        prev = node
        node = tmp
    }

    newHead := prev

    var compareHead *ListNode
    if n%2 == 0 {
        compareHead = node
    } else {
        compareHead = node.Next
    }

    for compareHead != nil {
        if newHead.Val != compareHead.Val {
            return false
        }
        newHead = newHead.Next
        compareHead = compareHead.Next
    }
    return true
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	for middle := head; ; {
		if head.Next != nil {
			middle, head = middle.Next, head.Next
		}
		if head = head.Next; head == nil {
			return middle
		}
	}
}



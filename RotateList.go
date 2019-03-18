package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}

	tail := head
	total := 0
	for tail.Next != nil {
		total++
		tail = tail.Next
	}
	total++

	k = k % total
	if k == 0 {
		return head
	}

	c := total - k
	nTail := head
	for i:=0; i < c-1; i ++ {
		nTail = nTail.Next
	}

	tail.Next = head
	head = nTail.Next
	nTail.Next = nil

	return head
}
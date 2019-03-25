package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	s := &ListNode{}
	s.Next = head

	t := s
	for t.Next != nil && t.Next.Next != nil {
		a := t.Next
		b := a.Next
		t.Next, b.Next, a.Next = b, a, b.Next
		t = a
	}

	return s.Next
}